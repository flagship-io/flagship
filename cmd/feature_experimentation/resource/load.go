/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/d5/tengo/v2"
	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	resourceFile string
	outputFile   string
)

type Data interface {
	Save(data string) ([]byte, error)
	Delete(id string) error
}

type ProjectData struct {
	*models.Project
}

type ResourceData struct {
	Id string `json:"id"`
}

func (f *ProjectData) Save(data string) ([]byte, error) {
	return http_request.ProjectRequester.HTTPCreateProject([]byte(data))
}

func (f *ProjectData) Delete(id string) error {
	return http_request.ProjectRequester.HTTPDeleteProject(id)
}

type CampaignData struct {
	Id              string               `json:"id,omitempty"`
	ProjectId       string               `json:"project_id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Type            string               `json:"type,omitempty"`
	VariationGroups []VariationGroupData `json:"variation_groups"`
}

func (f *CampaignData) Save(data string) ([]byte, error) {
	return http_request.CampaignRequester.HTTPCreateCampaign(data)
}

func (f *CampaignData) Delete(id string) error {
	return http_request.CampaignRequester.HTTPDeleteCampaign(id)
}

type FlagData struct {
	*models.Flag
}

func (f *FlagData) Save(data string) ([]byte, error) {
	return http_request.FlagRequester.HTTPCreateFlag(data)
}

func (f *FlagData) Delete(id string) error {
	return http_request.FlagRequester.HTTPDeleteFlag(id)
}

type GoalData struct {
	*models.Goal
}

func (f *GoalData) Save(data string) ([]byte, error) {
	return http_request.GoalRequester.HTTPCreateGoal(data)
}

func (f *GoalData) Delete(id string) error {
	return http_request.GoalRequester.HTTPDeleteGoal(id)
}

type TargetingKeysData struct {
	*models.TargetingKey
}

func (f *TargetingKeysData) Save(data string) ([]byte, error) {
	return http_request.TargetingKeyRequester.HTTPCreateTargetingKey(data)
}

func (f *TargetingKeysData) Delete(id string) error {
	return http_request.TargetingKeyRequester.HTTPDeleteTargetingKey(id)
}

type VariationGroupData struct {
	*models.VariationGroup
}

type VariationData struct {
	*models.Variation
}

type ResourceType int

const (
	Project ResourceType = iota
	Flag
	TargetingKey
	Goal
	Campaign
	VariationGroup
	Variation
)

var resourceTypeMap = map[string]ResourceType{
	"project":         Project,
	"flag":            Flag,
	"targeting_key":   TargetingKey,
	"goal":            Goal,
	"campaign":        Campaign,
	"variation_group": VariationGroup,
	"variation":       Variation,
}

type Resource struct {
	Name             ResourceType
	Data             Data
	ResourceVariable string
	Method           string
}

var cred common.RequestConfig

func Init(credL common.RequestConfig) {
	cred = credL
}

type ResourceCmdStruct struct {
	Name string `json:"name,omitempty"`
	Data string `json:"data,omitempty"`
}

func UnmarshalConfig(filePath string) ([]Resource, error) {
	var config struct {
		Resources []struct {
			Name             string
			Data             json.RawMessage
			ResourceVariable string
			Method           string
		}
	}

	bytes, err := os.ReadFile(resourceFile)

	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	var resources []Resource
	for _, r := range config.Resources {
		name, ok := resourceTypeMap[r.Name]
		if !ok {
			return nil, fmt.Errorf("invalid resource name: %s", r.Name)
		}

		var data Data = nil
		var err error = nil

		switch name {

		case Project:
			projectData := ProjectData{}
			err = json.Unmarshal(r.Data, &projectData)
			data = &projectData

		case Flag:
			flagData := FlagData{}
			err = json.Unmarshal(r.Data, &flagData)
			data = &flagData

		case TargetingKey:
			targetingKeyData := TargetingKeysData{}
			err = json.Unmarshal(r.Data, &targetingKeyData)
			data = &targetingKeyData

		case Campaign:
			campaignData := CampaignData{}
			err = json.Unmarshal(r.Data, &campaignData)
			data = &campaignData

		case Goal:
			goalData := GoalData{}
			err = json.Unmarshal(r.Data, &goalData)
			data = &goalData
		}

		if err != nil {
			return nil, err
		}

		resources = append(resources, Resource{Name: name, Data: data, ResourceVariable: r.ResourceVariable, Method: r.Method})
	}

	return resources, nil
}

var gResources []Resource

// LoadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load [--file=<file>]",
	Short: "Load your resources",
	Long:  `Load your resources`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonBytes := ScriptResource(cmd, gResources)
		if outputFile != "" {
			os.WriteFile(outputFile, jsonBytes, os.ModePerm)
			fmt.Fprintf(cmd.OutOrStdout(), "File created at %s\n", outputFile)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s", string(jsonBytes))
	},
}

func init() {
	cobra.OnInitialize(initResource)

	loadCmd.Flags().StringVarP(&resourceFile, "file", "", "", "resource file that contains your resource")

	if err := loadCmd.MarkFlagRequired("file"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	loadCmd.Flags().StringVarP(&outputFile, "output-file", "", "", "result of the command that contains all resource informations")

	ResourceCmd.AddCommand(loadCmd)
}

func initResource() {

	// Use config file from the flag.
	var err error
	if resourceFile != "" {
		gResources, err = UnmarshalConfig(resourceFile)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
	}
}

func ScriptResource(cmd *cobra.Command, resources []Resource) []byte {

	resourceVariables := make(map[string]interface{})
	var loadResultJSON []string
	var loadResultOutputFile []ResourceCmdStruct

	for _, resource := range resources {
		var response []byte
		var resultOutputFile ResourceCmdStruct
		var resourceData map[string]interface{}
		var responseData interface{}
		var url = ""
		var resourceName = ""
		const color = "\033[0;33m"
		const colorNone = "\033[0m"

		data, err := json.Marshal(resource.Data)
		if err != nil {
			fmt.Printf("error occurred marshal data: %v\n", err)
		}

		var httpMethod string = "POST"

		if resource.Method == "delete" {
			httpMethod = "DELETE"
		}

		switch resource.Name {
		case Project:
			url = "/projects"
			resourceName = "Project"
		case Flag:
			url = "/flags"
			resourceName = "Flag"
		case TargetingKey:
			url = "/targeting_keys"
			resourceName = "Targeting Key"
		case Goal:
			url = "/goals"
			resourceName = "Goal"
		case VariationGroup:
			url = "/variation_groups"
			resourceName = "Variation Group"
		case Variation:
			url = "/variations"
			resourceName = "Variation"
		case Campaign:
			url = "/campaigns"
			resourceName = "Campaign"
		}

		err = json.Unmarshal(data, &resourceData)

		if err != nil {
			fmt.Printf("error occurred unmarshall resourceData: %v\n", err)
		}

		for k, vInterface := range resourceData {
			v, ok := vInterface.(string)
			if ok {
				if strings.Contains(v, "$") {
					vTrim := strings.Trim(v, "$")
					for k_, variable := range resourceVariables {
						script, _ := tengo.Eval(context.Background(), vTrim, map[string]interface{}{
							k_: variable,
						})
						if script == nil {
							continue
						}
						resourceData[k] = script.(string)
					}
				}
			}
		}

		dataResource, err := json.Marshal(resourceData)
		if err != nil {
			log.Fatalf("error occurred http call: %v\n", err)
		}

		if httpMethod == "POST" {
			if resource.Name == Project || resource.Name == TargetingKey || resource.Name == Flag {
				response, err = common.HTTPRequest[ResourceData](httpMethod, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+cred.AccountID+url, dataResource)
			}

			if resource.Name == Goal || resource.Name == Campaign {
				response, err = common.HTTPRequest[ResourceData](httpMethod, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+cred.AccountID+"/account_environments/"+cred.AccountEnvironmentID+url, dataResource)
			}

			resultOutputFile = ResourceCmdStruct{
				Name: resourceName,
				Data: string(response),
			}
			loadResultOutputFile = append(loadResultOutputFile, resultOutputFile)

		}

		if httpMethod == "DELETE" {
			if resource.Name == Project || resource.Name == TargetingKey || resource.Name == Flag {
				response, err = common.HTTPRequest[ResourceData](httpMethod, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+cred.AccountID+url+"/"+fmt.Sprintf("%v", resourceData["id"]), nil)
			}

			if resource.Name == Goal || resource.Name == Campaign {
				response, err = common.HTTPRequest[ResourceData](httpMethod, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+cred.AccountID+"/account_environments/"+cred.AccountEnvironmentID+url+"/"+fmt.Sprintf("%v", resourceData["id"]), nil)
			}

			if err == nil && viper.GetString("output_format") != "json" {
				response = []byte("The id: " + fmt.Sprintf("%v", resourceData["id"]) + " deleted successfully")
			}
		}

		if err != nil {
			log.Fatalf("error occurred http call: %v\n", err)
		}

		if viper.GetString("output_format") != "json" {
			fmt.Fprintf(cmd.OutOrStdout(), "%s - %s: %s %s\n", color, resourceName, colorNone, string(response))
		}

		if httpMethod != "DELETE" {
			err = json.Unmarshal(response, &responseData)

			if err != nil {
				fmt.Printf("error occurred unmarshal responseData: %v\n", err)
			}

			if responseData == nil {
				fmt.Println("error occurred not response data: " + string(response))
				continue
			}

			resourceVariables[resource.ResourceVariable] = responseData
		}

		loadResultJSON = append(loadResultJSON, string(response))
	}

	var jsonBytes []byte
	var jsonString any

	if outputFile != "" {
		jsonString = loadResultOutputFile
	} else {
		jsonString = loadResultJSON
	}

	jsonBytes, err := json.Marshal(jsonString)

	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}

	return jsonBytes
}
