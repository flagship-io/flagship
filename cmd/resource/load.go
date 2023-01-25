/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package resource

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

func (f ProjectData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateProject(data)
}

type CampaignData struct {
	Id          string
	ProjectId   string
	Name        string
	Description string
	Type        string
	//VariationGroups []VariationGroupData
}

func (f CampaignData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateCampaign(data)
}

type FlagData struct {
	Id               string
	Name             string
	Type             string
	Description      string
	Source           string
	DefaultValue     string
	PredefinedValues []string
}

func (f FlagData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateFlag(data)
}

type GoalData struct {
	Id       string
	Label    string
	Type     string
	Operator string
	Value    string
}

func (f GoalData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateGoal(data)
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

type TargetingKeysData struct {
	Id          string
	Name        string
	Type        string
	Description string
}

func (f TargetingKeysData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateTargetingKey(data)
}

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
	Name ResourceType
	Data Data
}

func UnmarshalConfig(filePath string) ([]Resource, error) {
	var config struct {
		Resources []struct {
			Name string
			Data json.RawMessage
		}
	}

	bytes, err := os.ReadFile(resourceFile)

	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		fmt.Println("here")
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

		resources = append(resources, Resource{Name: name, Data: data})
	}

	//flag := resources[1].Data.(ProjectData).Name
	//fmt.Println(flag)
	return resources, nil
}

func loadResources(resources []Resource) (string, error) {

	for _, resource := range resources {
		var url = ""
		data, err := json.Marshal(resource.Data)
		if err != nil {
			return "", err
		}
		fmt.Println(string(data))

		switch resource.Name {
		case Project:
			url = "/project"
		}
		httprequest.HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+url, data)

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
		loadResources(gResources)
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
