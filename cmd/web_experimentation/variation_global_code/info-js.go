/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"log"
	"strconv"

	"github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoJSCmd represents info-js command
var infoJSCmd = &cobra.Command{
	Use:   "info-js [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global js code info",
	Long:  `Get variation global js code info `,
	Run: func(cmd *cobra.Command, args []string) {
		var modif web_experimentation.Modification

		campaignID, err := strconv.Atoi(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		variationID, err := strconv.Atoi(VariationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		body, err := httprequest.ModificationRequester.HTTPListModification(campaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range body {
			if modification.VariationID == variationID && modification.Type == "customScriptNew" && modification.Selector == "" {
				modif = modification
			}
		}

		utils.FormatItem([]string{"Id", "Name", "Type", "VariationID", "Selector", "Engine", "Value"}, modif, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	infoJSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")

	if err := infoJSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	infoJSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")

	if err := infoJSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	VariationGlobalCodeCmd.AddCommand(infoJSCmd)
}
