/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle [-i <project-id> | --id=<project-id>] [-s <status> | --status=<status>]",
	Short: "Toggle a project",
	Long:  `Toggle a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(ProjectStatus == "active" || ProjectStatus == "paused" || ProjectStatus == "interrupted") {
			fmt.Fprintln(cmd.OutOrStdout(), "Status can only have 3 values: active or paused or interrupted")
			return
		}

		err := httprequest.HTTPToggleProject(ProjectId, ProjectStatus)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "project set to %v\n", ProjectStatus)
	},
}

func init() {

	toggleCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to toggle")
	toggleCmd.Flags().StringVarP(&ProjectStatus, "status", "s", "", "status you want to set to the project. Only 3 values are possible: active, paused and interrupted")

	if err := toggleCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ProjectCmd.AddCommand(toggleCmd)
}
