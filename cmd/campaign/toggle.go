/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle [-i <campaign-id> | --campaign-id=<campaign-id>] [-s <status> | --status=<status>]",
	Short: "Toggle a campaign",
	Long:  `Toggle a campaign of your account`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(Status == "active" || Status == "paused" || Status == "interrupted") {
			log.Println("Status can only have 3 values : active or paused or interrupted")
		} else {
			err := httprequest.HTTPToggleCampaign(CampaignID, Status)
			if err != nil {
				log.Fatalf("error occured: %v", err)
			}
			log.Printf("campaign status set to %s", Status)
		}

	},
}

func init() {

	toggleCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to toggle")
	toggleCmd.Flags().StringVarP(&Status, "status", "s", "", "status you want set to the campaign. Only 3 value are possible: active, paused and interrupted")

	if err := toggleCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(toggleCmd)
}
