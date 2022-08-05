/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [--data-raw=<data-raw>]",
	Short: "Create a campaign",
	Long:  `Create a campaign in your project`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateCampaign(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("campaign created: %s", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your campaign, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(createCmd)
}
