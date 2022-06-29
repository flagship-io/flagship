/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name string
)

func createProject(project string) string {
	return "create project " + project + "in the account id " + viper.GetViper().GetString("account_id")
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(createProject(name))
	},
}

func init() {

	createCmd.Flags().StringVarP(&name, "name", "n", "", "the name")

	if err := createCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	ProjectCmd.AddCommand(createCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
