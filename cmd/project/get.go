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
	Use:   "get",
	Short: "this get project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"))
	},
}

func init() {

	getCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "get project by project id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	ProjectCmd.AddCommand(getCmd)

}
