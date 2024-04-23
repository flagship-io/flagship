/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package element_modification_code

import (
	"fmt"
	"log"
	"strconv"

	"github.com/flagship-io/flagship/utils/config"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <modification-id> | --id <modification-id>] [--campaign-id <campaign-id>]",
	Short: "Get element modification code",
	Long:  `Get element modification code`,
	Run: func(cmd *cobra.Command, args []string) {
		var code string
		var selector string
		var variationID int

		campaignID, err := strconv.Atoi(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		modificationID, err := strconv.Atoi(ModificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		body, err := httprequest.ModificationRequester.HTTPGetModificationByID(campaignID, modificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range body {
			if modification.Type == "customScriptNew" && modification.Selector != "" {
				code = modification.Value
				selector = modification.Selector
				variationID = modification.VariationID
			}
		}

		if CreateFile {
			fileCode := config.AddHeaderSelectorComment(selector, code)
			elementModificationCodeDir := config.ElementModificationCodeDirectory(viper.GetString("working_dir"), httprequest.CampaignGlobalCodeRequester.AccountID, CampaignID, strconv.Itoa(variationID), ModificationID, selector, fileCode)
			fmt.Fprintln(cmd.OutOrStdout(), "Element code file generated successfully: ", elementModificationCodeDir)
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(code))
	},
}

func init() {
	getCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the campaign")

	if err := getCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getCmd.Flags().StringVarP(&ModificationID, "id", "i", "", "id of the global code modification you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	getCmd.Flags().BoolVarP(&CreateFile, "create-file", "", false, "create a file that contains campaign global code")

	ElementModificationCodeCmd.AddCommand(getCmd)
}
