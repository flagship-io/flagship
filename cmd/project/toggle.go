/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "this toggle project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !(ProjectStatus == "active" || ProjectStatus == "paused" || ProjectStatus == "interrupted") {
			fmt.Println("Status can only have 3 values: active or paused or interrupted")
			return
		}

		err := httprequest.HTTPToggleProject(ProjectId, ProjectStatus)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("project set to %v", ProjectStatus)
	},
}

func init() {

	toggleCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "the project id")
	toggleCmd.Flags().StringVarP(&ProjectStatus, "status", "s", "", "the project status")

	if err := toggleCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}

	ProjectCmd.AddCommand(toggleCmd)
}
