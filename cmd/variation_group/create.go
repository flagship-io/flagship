/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateVariationGroup(CampaignID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("variation group created: %s", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&CampaignID, "campaign-id", "i", "", "the campaign id")

	if err := createCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "the data")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationGroupCmd.AddCommand(createCmd)
}
