/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

var jsCode string
var jsFilePath string

// pushJsCmd represents push command
var pushJSCmd = &cobra.Command{
	Use:   "push-js [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "push variation global js code",
	Long:  `push variation global js code`,
	Run: func(cmd *cobra.Command, args []string) {
		var modificationId int
		var codeByte []byte

		if !utils.CheckSingleFlag(jsFilePath != "", jsCode != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (file, code)")
		}

		campaignID, err := strconv.Atoi(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		variationID, err := strconv.Atoi(VariationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		modifList, err := httprequest.ModificationRequester.HTTPListModification(campaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range modifList {
			if modification.VariationID == variationID && modification.Type == "customScriptNew" && modification.Selector == "" {
				modificationId = modification.Id
			}
		}

		if jsFilePath != "" {
			fileContent, err := os.ReadFile(jsFilePath)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			codeByte = fileContent
		}

		if jsCode != "" {
			codeByte = []byte(jsCode)
		}

		if modificationId == 0 {
			modificationToPush := web_experimentation.ModificationCodeCreateStruct{
				InputType:   "modification",
				Name:        "",
				Value:       string(codeByte),
				Selector:    "",
				Type:        "customScriptNew",
				Engine:      "",
				VariationID: variationID,
			}

			body, err := httprequest.ModificationRequester.HTTPCreateModification(campaignID, modificationToPush)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			fmt.Fprintln(cmd.OutOrStdout(), string(body))
			return
		}

		modificationToPush := web_experimentation.ModificationCodeEditStruct{
			InputType: "modification",
			Value:     string(codeByte),
		}

		body, err := httprequest.ModificationRequester.HTTPEditModification(campaignID, modificationId, modificationToPush)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(body))
	},
}

func init() {
	pushJSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")
	if err := pushJSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushJSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")
	if err := pushJSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushJSCmd.Flags().StringVarP(&jsCode, "code", "c", "", "new code to push in the variation")
	pushJSCmd.Flags().StringVarP(&jsFilePath, "file", "", "", "file that contains new code to push in the variation")

	VariationGlobalCodeCmd.AddCommand(pushJSCmd)
}
