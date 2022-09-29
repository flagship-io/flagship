/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package campaign

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle [-i <campaign-id> | --id=<campaign-id>] [-s <status> | --status=<status>]",
	Short: "Toggle a campaign",
	Long:  `Toggle a campaign in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(Status == "active" || Status == "paused" || Status == "interrupted") {
			fmt.Fprintln(cmd.OutOrStdout(), "Status can only have 3 values : active or paused or interrupted")
		} else {
			err := httprequest.HTTPToggleCampaign(CampaignID, Status)
			if err != nil {
				log.Fatalf("error occured: %v", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "campaign status set to %s\n", Status)
		}

	},
}

func init() {

	toggleCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to toggle")
	toggleCmd.Flags().StringVarP(&Status, "status", "s", "", "status you want set to the campaign. Only 3 values are possible: active, paused and interrupted")

	if err := toggleCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(toggleCmd)
}
