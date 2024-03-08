/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [--campaign-id=<campaign-id>] [--variation-group-id=<variation-group-id>] [-i <variation-id> | --id=<variation-id>] [-d <data-raw> | --data-raw=<data-raw>]",
	Short: "Edit a variation",
	Long:  `Edit a variation in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditVariation(CampaignID, VariationGroupID, VariationID, DataRaw)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the variation you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your variation, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	VariationCmd.AddCommand(editCmd)
}
