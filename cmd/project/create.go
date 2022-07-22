/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpCreateProject(ProjectName)
	},
}

func init() {

	createCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "the name")

	if err := createCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}

	ProjectCmd.AddCommand(createCmd)

}
