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

var cssCode string
var cssFilePath string

// pushCSSCmd represents push command
var pushCSSCmd = &cobra.Command{
	Use:   "push-css [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "push variation global css code",
	Long:  `push variation global css code`,
	Run: func(cmd *cobra.Command, args []string) {
		var modificationId int
		var codeByte []byte

		if !utils.CheckSingleFlag(cssFilePath != "", cssCode != "") {
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
			if modification.VariationID == variationID && modification.Type == "addCSS" && modification.Selector == "" {
				modificationId = modification.Id
			}
		}

		if cssFilePath != "" {
			fileContent, err := os.ReadFile(cssFilePath)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			codeByte = fileContent
		}

		if cssCode != "" {
			codeByte = []byte(cssCode)
		}

		if modificationId == 0 {
			modificationToPush := web_experimentation.ModificationCodeCreateStruct{
				InputType:   "modification",
				Name:        "",
				Value:       string(codeByte),
				Selector:    "",
				Type:        "addCSS",
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
	pushCSSCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the global code campaign you want to display")
	if err := pushCSSCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCSSCmd.Flags().StringVarP(&VariationID, "id", "i", "", "id of the global code variation you want to display")
	if err := pushCSSCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCSSCmd.Flags().StringVarP(&cssCode, "code", "c", "", "new code to push in the variation")
	pushCSSCmd.Flags().StringVarP(&cssFilePath, "file", "", "", "file that contains new code to push in the variation")

	VariationGlobalCodeCmd.AddCommand(pushCSSCmd)
}
