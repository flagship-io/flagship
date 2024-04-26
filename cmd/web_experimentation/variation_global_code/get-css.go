/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"fmt"
	"log"
	"strconv"

	"github.com/flagship-io/flagship/utils/config"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCSSCmd represents get command
var getCSSCmd = &cobra.Command{
	Use:   "get-css [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global css code",
	Long:  `Get variation global css code`,
	Run: func(cmd *cobra.Command, args []string) {
		var cssCode string

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
				cssCode = modification.Value
			}
		}

		if CreateFile {
			campaignCodeDir, err := config.VariationGlobalCodeDirectoryCSS(viper.GetString("working_dir"), httprequest.CampaignGlobalCodeRequester.AccountID, CampaignID, VariationID, cssCode)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Variation code file generated successfully: ", campaignCodeDir)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), cssCode)
	},
}

func init() {
	getCSSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")

	if err := getCSSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getCSSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")

	if err := getCSSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getCSSCmd.Flags().BoolVarP(&CreateFile, "create-file", "", false, "create a file that contains campaign global code")

	VariationGlobalCodeCmd.AddCommand(getCSSCmd)
}
