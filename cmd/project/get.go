/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nameGet string
)

func getProject(project string) string {
	return "get project " + project
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "get",
	Short: "this get project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getProject(nameGet))
	},
}

func init() {

	editCmd.Flags().StringVarP(&nameGet, "name", "n", "", "the name")

	if err := editCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(editCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
