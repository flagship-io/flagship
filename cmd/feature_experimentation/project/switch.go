/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request/feature_experimentation"
	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch [-i <project-id> | --id=<project-id>] [-s <status> | --status=<status>]",
	Short: "switch a project state",
	Long:  `switch a project state in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(ProjectStatus == "active" || ProjectStatus == "paused" || ProjectStatus == "interrupted") {
			fmt.Fprintln(cmd.OutOrStdout(), "Status can only have 3 values: active or paused or interrupted")
			return
		}

		err := httprequest.HTTPSwitchProject(ProjectId, ProjectStatus)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "project status set to %v\n", ProjectStatus)
	},
}

func init() {

	switchCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to switch state")
	switchCmd.Flags().StringVarP(&ProjectStatus, "status", "s", "", "status you want to set to the project. Only 3 values are possible: active, paused and interrupted")

	if err := switchCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := switchCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ProjectCmd.AddCommand(switchCmd)
}
