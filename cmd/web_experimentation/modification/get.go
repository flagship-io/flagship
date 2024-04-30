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

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <modification-id> | --id <modification-id>] [--campaign-id <campaign-id>]",
	Short: "Get a modification",
	Long:  `Get a modification`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.ModificationRequester.HTTPGetModification(CampaignID, ModificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name", "Type", "VariationID", "Selector", "Engine", "Value"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	getCmd.Flags().IntVarP(&ModificationID, "id", "i", 0, "id of the modification you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	ModificationCmd.AddCommand(getCmd)
}
