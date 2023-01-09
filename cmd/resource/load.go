/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package resource

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

type Data interface {
	Save(data string) ([]byte, error)
}

type ProjectData struct {
	Id   string
	Name string
}

func (f ProjectData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateProject(data)
}

type FlagData struct {
	Id          string
	Name        string
	Type        string
	Description string
	Source      string
}

func (f FlagData) Save(data string) ([]byte, error) {
	return httprequest.HTTPCreateFlag(data)
}

// define structs for other resource types

type ResourceType int

const (
	Project ResourceType = iota
	Flag
	TargetingKey
	Goal
)

var resourceTypeMap = map[string]ResourceType{
	"project":       Project,
	"flag":          Flag,
	"targeting_key": TargetingKey,
	"goal":          Goal,
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

			fmt.Println(data)

		//data = &ProjectData{}
		case Flag:
			flagData := FlagData{}
			err = json.Unmarshal(r.Data, &flagData)
			data = flagData
			fmt.Println(data)
		}

		if err != nil {
			return nil, err
		}

		resources = append(resources, Resource{Name: name, Data: data})
	}

	flag := resources[1].Data.(FlagData).Name
	fmt.Println(flag)
	return resources, nil
}

func loadResources(resources []Resource) ([]Resource, error) {

	for _, resource := range resources {
		data, err := json.Marshal(resource.Data)
		if err != nil {
			return nil, err
		}
		resource.Data.Save()
	}
}

var gResources []Resource

// LoadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load [--file=<file>]",
	Short: "Load your resources",
	Long:  `Load your resources`,
	Run: func(cmd *cobra.Command, args []string) {

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
	if resourceFile != "" {
		_, err := UnmarshalConfig(resourceFile)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
	}
}
