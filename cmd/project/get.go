/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	projectIdGet string
)

func getProject(project string) string {
	return "get project " + project
}
func getAllProject() string {
	return "get All projects"
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "get",
	Short: "this get project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getAll, _ := cmd.Flags().GetBool("all")
		if getAll {
			fmt.Println(getAllProject())
		} else if projectIdGet == "" {
			fmt.Println("project id required")
		} else {
			fmt.Println(getProject(projectIdGet))
		}

	},
}

func init() {

	editCmd.Flags().StringVarP(&projectIdGet, "project_id", "i", "", "get project by project id")
	//editCmd.Flags().BoolVarP(&getAll, "all", "", true, "get all project")
	editCmd.Flags().BoolP("all", "", false, "get All Project")

	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(editCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
