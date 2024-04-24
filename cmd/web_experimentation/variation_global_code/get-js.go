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

// getJsCmd represents get command
var getJSCmd = &cobra.Command{
	Use:   "get-js [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Get variation global js code",
	Long:  `Get variation global js code`,
	Run: func(cmd *cobra.Command, args []string) {
		var jsCode string

		campaignID, err := strconv.Atoi(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		variationID, err := strconv.Atoi(VariationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		body, err := httprequest.ModificationRequester.HTTPGetModification(campaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range body {
			if modification.VariationID == variationID && modification.Type == "customScriptNew" && modification.Selector == "" {
				jsCode = modification.Value
			}
		}

		if CreateFile {
			variationCodeDir, err := config.VariationGlobalCodeDirectoryJS(viper.GetString("working_dir"), httprequest.CampaignGlobalCodeRequester.AccountID, CampaignID, VariationID, jsCode)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "Variation code file generated successfully: ", variationCodeDir)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), jsCode)
	},
}

func init() {
	getJSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")

	if err := getJSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getJSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")

	if err := getJSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getJSCmd.Flags().BoolVarP(&CreateFile, "create-file", "", false, "create a file that contains campaign global code")

	VariationGlobalCodeCmd.AddCommand(getJSCmd)
}
