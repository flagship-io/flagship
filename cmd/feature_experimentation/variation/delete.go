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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [--campaign-id=<campaign-id>] [--variation-group-id=<variation-group-id>] [-i <variation-id> | --id=<variation-id>]",
	Short: "Delete a variation",
	Long:  `Delete a variation in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.VariationFERequester.HTTPDeleteVariation(CampaignID, VariationGroupID, VariationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Variation deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the variation you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Fprintf(deleteCmd.OutOrStderr(), "error occurred: %s", err)
	}
	VariationCmd.AddCommand(deleteCmd)
}
