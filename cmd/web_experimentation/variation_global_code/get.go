/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ModificationGlobalCode struct {
	JS  string `json:"js,omitempty"`
	CSS string `json:"css,omitempty"`
}

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global code",
	Long:  `Get variation global code`,
	Run: func(cmd *cobra.Command, args []string) {
		var modificationResp ModificationGlobalCode
		body, err := httprequest.ModificationRequester.HTTPGetModification(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range body {
			if modification.VariationID == VariationID && modification.Type == "customScriptNew" {
				modificationResp.JS = modification.Value
			}
			if modification.VariationID == VariationID && modification.Type == "addCSS" {
				modificationResp.CSS = modification.Value
			}
		}

		utils.FormatItem([]string{"JS", "CSS"}, modificationResp, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	VariationGlobalCodeCmd.AddCommand(getCmd)
}
