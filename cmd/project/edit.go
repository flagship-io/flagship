/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package project

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-i <project-id> | --id=<project-id>] [-n <name> | --name=<name>]",
	Short: "Edit a project",
	Long:  `Edit a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPEditProject(ProjectId, ProjectName)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Project updated")
	},
}

func init() {

	editCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to edit")

	editCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "name you want to set for the project")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	ProjectCmd.AddCommand(editCmd)

}
