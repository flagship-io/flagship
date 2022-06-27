/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nameEdit string
)

func editProject(project string) string {
	return "edit project " + project
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

	getCmd.Flags().StringVarP(&nameEdit, "name", "n", "", "the name")

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
