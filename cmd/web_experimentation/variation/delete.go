/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <variation-id> | --id=<variation-id>] [--campaign-id <campaign-id>]",
	Short: "Delete a variation",
	Long:  `Delete a variation`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.VariationWERequester.HTTPDeleteVariation(CampaignID, VariationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Variation deleted")

	},
}

func init() {
	deleteCmd.Flags().IntVarP(&VariationID, "id", "i", 0, "id of the variation you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	VariationCmd.AddCommand(deleteCmd)
}
