/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package campaign

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship-cli/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-d <data-raw> | --data-raw=<data-raw>]",
	Short: "Create a campaign",
	Long:  `Create a campaign in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateCampaign(DataRaw)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your campaign, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	CampaignCmd.AddCommand(createCmd)
}
