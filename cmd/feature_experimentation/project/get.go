/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package project

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [-i <project-id> | --id=<project-id>]",
	Short: "Get a project",
	Long:  `Get a project`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.ProjectRequester.HTTPGetProject(ProjectId)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {

	getCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "id of the project you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	ProjectCmd.AddCommand(getCmd)

}
