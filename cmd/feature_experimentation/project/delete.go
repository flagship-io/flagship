/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <project-id> | --id=<project-id>]",
	Short: "Delete a project",
	Long:  `Delete a project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.ProjectRequester.HTTPDeleteProject(ProjectId)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Project deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ProjectCmd.AddCommand(deleteCmd)
}
