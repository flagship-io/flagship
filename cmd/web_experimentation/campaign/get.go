/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <test-id> | --id <test-id>]",
	Short: "Get a test",
	Long:  `Get a test in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.CampaignWERequester.HTTPGetCampaign(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name", "Description", "Type", "State", "Url"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	getCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the test you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	CampaignCmd.AddCommand(getCmd)
}
