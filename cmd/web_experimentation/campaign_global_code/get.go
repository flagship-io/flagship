/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign_global_code

import (
	"fmt"
	"log"

	"github.com/flagship-io/flagship/utils/config"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var campaignID string
var createFile bool

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

		if createFile {
			accountCodeDir := config.CampaignGlobalCodeDirectory(viper.GetString("working_dir"), httprequest.CampaignGlobalCodeRequester.AccountID, campaignID, body)
			fmt.Fprintln(cmd.OutOrStdout(), "Campaign code file generated successfully: ", accountCodeDir)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), body)

	},
}

func init() {
	getCmd.Flags().StringVarP(&campaignID, "id", "i", "", "id of the global code campaign you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	getCmd.Flags().BoolVarP(&createFile, "create-file", "", false, "create a file that contains campaign global code")

	CampaignGlobalCodeCmd.AddCommand(getCmd)
}
