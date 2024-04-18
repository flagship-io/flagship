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
	JS          string `json:"js,omitempty"`
	CSS         string `json:"css,omitempty"`
	VariationID int    `json:"variation_id,omitempty"`
	CampaignID  int    `json:"campaign_id,omitempty"`
}

// getCmd represents get command
var GetCmd = &cobra.Command{
	Use:   "get [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global code",
	Long:  `Get variation global code`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := GetCodeFiles(VariationID, CampaignID)

		utils.FormatItem([]string{"VariationID", "CampaignID", "JS", "CSS"}, resp, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	GetCmd.Flags().IntVarP(&CampaignID, "campaign-id", "", 0, "id of the global code campaign you want to display")

	if err := GetCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	GetCmd.Flags().IntVarP(&VariationID, "id", "i", 0, "id of the global code vairation you want to display")

	if err := GetCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	VariationGlobalCodeCmd.AddCommand(GetCmd)
}

func GetCodeFiles(variationID, campaignID int) ModificationGlobalCode {
	var modificationResp ModificationGlobalCode
	body, err := httprequest.ModificationRequester.HTTPGetModification(campaignID)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	for _, modification := range body {
		if modification.VariationID == variationID && modification.Type == "customScriptNew" {
			modificationResp.JS = modification.Value
		}
		if modification.VariationID == variationID && modification.Type == "addCSS" {
			modificationResp.CSS = modification.Value
		}
	}

	modificationResp.CampaignID = campaignID
	modificationResp.VariationID = variationID

	return modificationResp
}
