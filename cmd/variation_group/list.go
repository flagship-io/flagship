/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package variation_group

import (
	"log"

	"github.com/flagship-io/flagship-cli/utils"
	httprequest "github.com/flagship-io/flagship-cli/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--campaign-id=<campaign-id>]",
	Short: "List all variation groups",
	Long:  `List all variation groups in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListVariationGroup(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	VariationGroupCmd.AddCommand(listCmd)
}
