/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

var (
	projectIdDelete string
)

func deleteProject(project string) string {
	return "delete project id: " + project
}

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(deleteProject(projectIdDelete))
		httprequest.HttpDeleteProject(projectIdDelete)
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&projectIdDelete, "project_id", "i", "", "the projectid")

	if err := deleteCmd.MarkFlagRequired("project_id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(deleteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
