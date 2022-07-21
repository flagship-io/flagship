/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpDeleteProject(ProjectId)
		fmt.Println("Project deleted successfully.")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "the project id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	ProjectCmd.AddCommand(deleteCmd)
}
