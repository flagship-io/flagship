/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign_global_code

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

var campaignID string

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <campaign-id> | --id <campaign-id>]",
	Short: "Get campaign global code",
	Long:  `Get campaign global code`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.CampaignGlobalCodeRequester.HTTPGetCampaignGlobalCode(campaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), body)

	},
}

func init() {
	getCmd.Flags().StringVarP(&campaignID, "id", "i", "", "id of the global code campaign you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	CampaignGlobalCodeCmd.AddCommand(getCmd)
}
