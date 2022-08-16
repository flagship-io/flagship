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
	Use:   "delete",
	Short: "this delete project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Project deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "the project id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	ProjectCmd.AddCommand(deleteCmd)
}
