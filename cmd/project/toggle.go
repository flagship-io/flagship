/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

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
			fmt.Println("Status can only have 3 value: active or paused or interrupted")
		} else {
			httprequest.HttpToggleProject(ProjectId, ProjectStatus)
		}
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
