/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package variation_group

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [--campaign-id=<campaign-id>] [-i <variation-group-id> | --id <variation-group-id>]",
	Short: "Delete a variation group",
	Long:  `Delete a variation group in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteVariationGroup(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Variation group deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "id of the variation group you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	VariationGroupCmd.AddCommand(deleteCmd)
}
