/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign

import (
	"encoding/json"
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
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
		jsonBody, err := json.Marshal(body)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Printf("%s", string(jsonBody))
	},
}

func init() {
	getCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the test you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	CampaignCmd.AddCommand(getCmd)
}
