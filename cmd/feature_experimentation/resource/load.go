/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	featureexp "github.com/flagship-io/flagship/utils/http_request/feature_experimentation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Data interface {
	Save(data string) ([]byte, error)
}

type ProjectData struct {
	*models.Project
}

type ResourceData struct {
	Id string `json:"id"`
}

func (f *ProjectData) Save(data string) ([]byte, error) {
	return featureexp.HTTPCreateProject([]byte(data))
}

type CampaignData struct {
	Id              string               `json:"id,omitempty"`
	ProjectId       string               `json:"project_id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Type            string               `json:"type"`
	VariationGroups []VariationGroupData `json:"variation_groups"`
}

func (f *CampaignData) Save(data string) ([]byte, error) {
	return featureexp.HTTPCreateCampaign(data)
}

type FlagData struct {
	*models.Flag
}

func (f *FlagData) Save(data string) ([]byte, error) {
	return featureexp.HTTPCreateFlag(data)
}

type GoalData struct {
	*models.Goal
}

func (f *GoalData) Save(data string) ([]byte, error) {
	return featureexp.HTTPCreateGoal(data)
}

type TargetingKeysData struct {
	*models.TargetingKey
}

func (f *TargetingKeysData) Save(data string) ([]byte, error) {
	return featureexp.HTTPCreateTargetingKey(data)
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
}

func UnmarshalConfig(filePath string) ([]Resource, error) {
	var config struct {
		Resources []struct {
			Name             string
			Data             json.RawMessage
			ResourceVariable string
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

		resources = append(resources, Resource{Name: name, Data: data, ResourceVariable: r.ResourceVariable})
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
		ScriptResource(gResources)
	},
}

func init() {
	cobra.OnInitialize(initResource)

	loadCmd.Flags().StringVarP(&resourceFile, "file", "", "", "resource file that contains your resource")

	if err := loadCmd.MarkFlagRequired("file"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

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

func ScriptResource(resources []Resource) {

	resourceVariables := make(map[string]interface{})

	for _, resource := range resources {
		var response []byte
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

		if resource.Name == Project || resource.Name == TargetingKey || resource.Name == Flag {
			response, err = httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+url, dataResource)
		}

		if resource.Name == Goal || resource.Name == Campaign {
			response, err = httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+url, dataResource)
		}

		if err != nil {
			log.Fatalf("error occurred http call: %v\n", err)
		}

		fmt.Fprintf(os.Stdout, "%s - %s: %s %s\n", color, resourceName, colorNone, string(response))

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
}
