/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package modification

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--campaign-id <campaign-id>]",
	Short: "List all modifications",
	Long:  `List all modifications of a campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.ModificationRequester.HTTPListModification(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name", "Type", "VariationID", "Selector", "Engine", "Value"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	ModificationCmd.AddCommand(listCmd)
}
