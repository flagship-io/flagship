/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nameEdit      string
	projectIdEdit string
)

func editProject(name string) string {
	return "edit project that had id " + projectIdEdit + " with new name " + name
}

// createCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(editProject(nameEdit))
	},
}

func init() {

	getCmd.Flags().StringVarP(&projectIdEdit, "project_id", "i", "", "the project id")

	getCmd.Flags().StringVarP(&nameEdit, "name", "n", "", "the project new name")

	if err := getCmd.MarkFlagRequired("project_id"); err != nil {
		fmt.Println(err)
	}

	if err := getCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(getCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
