/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditVariationGroup(CampaignID, VariationGroupID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("variation group updated: %s", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "the campaign id")
	editCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "the variation group id")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "the data raw")

	if err := editCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	VariationGroupCmd.AddCommand(editCmd)
}
