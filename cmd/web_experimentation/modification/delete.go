/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package modification

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <modification-id> | --id=<modification-id>] [--campaign-id <campaign-id>]",
	Short: "Delete a modification",
	Long:  `Delete a modification`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.ModificationRequester.HTTPDeleteModification(CampaignID, ModificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Modification deleted")

	},
}

func init() {
	deleteCmd.Flags().IntVarP(&ModificationID, "id", "i", 0, "id of the modification you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ModificationCmd.AddCommand(deleteCmd)
}
