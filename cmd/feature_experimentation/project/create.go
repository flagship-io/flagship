/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package project

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-n <name> | --name=<name>]",
	Short: "Create a project",
	Long:  `Create a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		projectRequest := models.Project{
			Name: ProjectName,
		}

		projectRequestJSON, err := json.Marshal(projectRequest)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		body, err := httprequest.HTTPCreateProject(projectRequestJSON)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "name of the project you want to create")

	if err := createCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ProjectCmd.AddCommand(createCmd)
}
