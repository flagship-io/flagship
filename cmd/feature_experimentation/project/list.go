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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  `List all projects in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.ProjectRequester.HTTPListProject()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	ProjectCmd.AddCommand(listCmd)
}
