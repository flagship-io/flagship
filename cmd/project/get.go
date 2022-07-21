/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpGetProject(ProjectId)
	},
}

func init() {

	getCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "get project by project id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	ProjectCmd.AddCommand(getCmd)

}
