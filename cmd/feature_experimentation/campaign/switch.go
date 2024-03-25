/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// SwitchCmd represents the Switch command
var SwitchCmd = &cobra.Command{
	Use:   "switch [-i <campaign-id> | --id=<campaign-id>] [-s <status> | --status=<status>]",
	Short: "Switch a campaign state",
	Long:  `Switch a campaign state in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(Status == "active" || Status == "paused" || Status == "interrupted") {
			fmt.Fprintln(cmd.OutOrStdout(), "Status can only have 3 values: active or paused or interrupted")
		} else {
			err := httprequest.CampaignRequester.HTTPSwitchCampaign(CampaignID, Status)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "campaign status set to %s\n", Status)
		}

	},
}

func init() {

	SwitchCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to switch state")
	SwitchCmd.Flags().StringVarP(&Status, "status", "s", "", "status you want set to the campaign. Only 3 values are possible: active, paused and interrupted")

	if err := SwitchCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := SwitchCmd.MarkFlagRequired("status"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	CampaignCmd.AddCommand(SwitchCmd)
}
