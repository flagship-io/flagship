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
	Use:   "edit [--campaign-id=<campaign-id>] [-i <variation-group-id> | --id <variation-group-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Edit a variation group",
	Long:  `Edit a variation group in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditVariationGroup(CampaignID, VariationGroupID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("variation group updated: %s", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "id of the variation group you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your variation group, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	VariationGroupCmd.AddCommand(editCmd)
}
