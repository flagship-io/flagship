/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <project-id> | --id=<project-id>]",
	Short: "Delete a project",
	Long:  `Delete a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("project deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	ProjectCmd.AddCommand(deleteCmd)
}
