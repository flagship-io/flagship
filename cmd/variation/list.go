/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package variation

import (
	"log"

	"github.com/flagship-io/flagship-cli/utils"
	httprequest "github.com/flagship-io/flagship-cli/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--campaign-id=<campaign-id>] [--variation-group-id=<variation-group-id>]",
	Short: "List all variations",
	Long:  `List all variations in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListVariation(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Reference", "Allocation"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	VariationCmd.AddCommand(listCmd)
}
