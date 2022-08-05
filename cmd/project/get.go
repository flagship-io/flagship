/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [-i <project-id> | --id=<project-id>]",
	Short: "Get a project",
	Long:  `Get a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"))
	},
}

func init() {

	getCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	ProjectCmd.AddCommand(getCmd)

}
