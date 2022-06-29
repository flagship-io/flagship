/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	projectIdToggle string
	projectStatus   string
)

func toggleProject(project string) string {
	return "toggle project " + project + " with status " + projectStatus
}

// createCmd represents the create command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "this toggle project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !(projectStatus == "TRUE" || projectIdToggle == "FALSE") {
			fmt.Println("Status can have only 2 value TRUE or FALSE")
		} else {
			fmt.Println(toggleProject(projectIdToggle))
		}
	},
}

func init() {

	toggleCmd.Flags().StringVarP(&projectIdToggle, "project_id", "i", "", "the project id")
	toggleCmd.Flags().StringVarP(&projectStatus, "status", "s", "", "the project id")

	if err := toggleCmd.MarkFlagRequired("project_id"); err != nil {
		fmt.Println(err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(toggleCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
