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

// infoCSSCmd represents info-css command
var infoCSSCmd = &cobra.Command{
	Use:   "info-css [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global css code info",
	Long:  `Get variation global css code info `,
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
			if modification.VariationID == variationID && modification.Type == "addCSS" && modification.Selector == "" {
				modif = modification
			}
		}

		utils.FormatItem([]string{"Id", "Name", "Type", "VariationID", "Selector", "Engine", "Value"}, modif, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	infoCSSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")

	if err := infoCSSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	infoCSSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")

	if err := infoCSSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	VariationGlobalCodeCmd.AddCommand(infoCSSCmd)
}
