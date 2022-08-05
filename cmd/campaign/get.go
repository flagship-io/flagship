/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [-i <campaign-id> | --campaign-id=<campaign-id>]",
	Short: "Get a campaign",
	Long:  `Get a campaign of your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetCampaign(CampaignID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "ProjectID", "Name", "Description", "Type", "Status"}, body, viper.GetString("output_format"))
	},
}

func init() {

	getCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(getCmd)
}
