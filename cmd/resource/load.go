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
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Data interface {
	Save(data string) ([]byte, error)
}

type ProjectData struct {
	Id   string `json:",omitempty"`
	Name string `json:"name"`
}

type ResourceData struct {
	Id string `json:"id,omitempty"`
}

func (f ProjectData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateProject(data)
}

type CampaignData struct {
	Id          string `json:",omitempty"`
	ProjectId   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	//VariationGroups []VariationGroupData
}

func (f CampaignData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateCampaign(data)
}

type FlagData struct {
	Id               string   `json:",omitempty"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Source           string   `json:"source"`
	DefaultValue     string   `json:",omitempty"`
	PredefinedValues []string `json:",omitempty"`
}

func (f FlagData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateFlag(data)
}

type GoalData struct {
	Id       string `json:",omitempty"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

func (f GoalData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateGoal(data)
}

type TargetingKeysData struct {
	Id          string `json:",omitempty"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (f TargetingKeysData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateTargetingKey(data)
}

/* type VariationGroupData struct {
	Id            string
	Name          string
	Allocation    int
	Reference     bool
	Modifications interface{}
	Value         string
}

func (f VariationGroupData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateVariationGroup(campaignID, data)
}

type VariationData struct {
	Id        string
	Varations []VariationData
}

func (f VariationData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateVariation(campaignID, variationGroupID, data)
} */

// define structs for other resource types

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
			data = projectData
			//fmt.Println(data)

		//data = &ProjectData{}
		case Flag:
			flagData := FlagData{}
			err = json.Unmarshal(r.Data, &flagData)
			data = flagData
			//fmt.Println(data)

		case TargetingKey:
			targetingKeyData := TargetingKeysData{}
			err = json.Unmarshal(r.Data, &targetingKeyData)
			data = targetingKeyData
			//fmt.Println(data)

		case Campaign:
			campaignData := CampaignData{}
			err = json.Unmarshal(r.Data, &campaignData)
			data = campaignData

		case Goal:
			goalData := GoalData{}
			err = json.Unmarshal(r.Data, &goalData)
			data = goalData
			//fmt.Println(data)

			/* 		case VariationGroup:
			variationGroupData := VariationGroupData{}
			err = json.Unmarshal(r.Data, &variationGroupData)
			data = variationGroupData
			fmt.Println(data) */

		}

		if err != nil {
			return nil, err
		}

		resources = append(resources, Resource{Name: name, Data: data, ResourceVariable: r.ResourceVariable})
	}

	//flag := resources[1].Data.(ProjectData).Name
	//fmt.Println(flag)
	return resources, nil
}

func loadResources(resources []Resource) (string, error) {

	for _, resource := range resources {
		var url = ""
		var resp []byte
		data, err := json.Marshal(resource.Data)
		if err != nil {
			return "", err
		}

		switch resource.Name {
		case Project:
			url = "/projects"
		case Flag:
			url = "/flags"
		case TargetingKey:
			url = "/targeting_keys"
		case Goal:
			url = "/goals"
		case VariationGroup:
			url = "/variable_groups"
		case Variation:
			url = "/variations"
		case Campaign:
			url = "/campaigns"
		}

		if resource.Name == Project || resource.Name == TargetingKey || resource.Name == Flag {
			resp, err = httprequest.HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+url, data)
		}

		if resource.Name == Goal || resource.Name == Campaign {
			resp, err = httprequest.HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+url, data)
		}

		if err != nil {
			return "", err
		}

		log.Println(string(resp))

	}
	return "done", nil
}

var gResources []Resource

// LoadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load [--file=<file>]",
	Short: "Load your resources",
	Long:  `Load your resources`,
	Run: func(cmd *cobra.Command, args []string) {
		/* 		res, err := loadResources(gResources)
		   		if err != nil {
		   			log.Fatalf("error occurred: %v", err)
		   		}
		   		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", res) */
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

/* func GetResourceVariableAndData(resources []Resource) ([]string, []interface{}) {

	var resourceVariables []string
	var data []interface{}

	for _, r := range resources {
		data = append(data, r.Data)
		resourceVariables = append(resourceVariables, r.ResourceVariable)
	}

	return resourceVariables, data
} */

func ScriptResource(resources []Resource) {

	//var resourceVariables map[string]interface{}
	resourceVariables := make(map[string]interface{})

	for _, resource := range resources {
		var resourceData map[string]string
		var response []byte
		var responseData ResourceData
		var url = ""

		data, err := json.Marshal(resource.Data)
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
		}
		//fmt.Println(string(data), url)

		switch resource.Name {
		case Project:
			url = "/projects"
		case Flag:
			url = "/flags"
		case TargetingKey:
			url = "/targeting_keys"
		case Goal:
			url = "/goals"
		case VariationGroup:
			url = "/variable_groups"
		case Variation:
			url = "/variations"
		case Campaign:
			url = "/campaigns"
		}

		err = json.Unmarshal(data, &resourceData)

		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
		}

		for k, v := range resourceData {
			if strings.Contains(v, "$") {
				vTrim := strings.Trim(v, "$")
				script := tengo.NewScript([]byte(vTrim))
				for k_, variable := range resourceVariables {
					err = script.Add(k_, variable)
					if err != nil {
						fmt.Printf("error occurred affectation: %v\n", err)
					}
					fmt.Println("hello", k_, variable)
					fmt.Println("hi", k, vTrim)
				}

				compiled, err := script.RunContext(context.Background())
				if err != nil {
					fmt.Printf("error occurred compiled: %v\n", err)
				}
				resourceData[k] = compiled.Get(vTrim).String()
				fmt.Println(compiled.Get(vTrim).String())
			}

		}

		fmt.Println(resource.ResourceVariable)
		fmt.Println(resourceData)

		//fmt.Println(responseData)
		//fmt.Println(resource)

		if resource.Name == Project || resource.Name == TargetingKey || resource.Name == Flag {
			response, err = httprequest.HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+url, data)
		}

		if resource.Name == Goal || resource.Name == Campaign {
			response, err = httprequest.HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+url, data)
		}

		if err != nil {
			log.Fatalf("error occurred: %v\n", err)
		}

		err = json.Unmarshal(response, &responseData)

		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
		}

		if responseData.Id == "" {
			fmt.Println("error occurred: " + string(response))
			continue
		}

		resourceVariables[resource.ResourceVariable] = responseData.Id
		fmt.Println(resourceVariables)
	}
}
