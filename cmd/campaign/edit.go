/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-i <campaign-id> | --id=<campaign-id>] [ -d <data-raw> | --data-raw=<data-raw>]",
	Short: "Edit a campaign",
	Long:  `Edit a campaign in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditCampaign(CampaignID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "campaign updated: %s\n", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your campaign, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(editCmd)
}
