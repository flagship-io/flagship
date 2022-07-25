/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPEditProject(ProjectId, ProjectName)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("project updated")
	},
}

func init() {

	editCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "the project id")

	editCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "the project new name")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	if err := editCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	ProjectCmd.AddCommand(editCmd)

}
