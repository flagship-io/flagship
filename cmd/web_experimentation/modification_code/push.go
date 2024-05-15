/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package modification_code

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

var code string
var filePath string
var selector string
var variationID string

// pushCmd represents get command
var pushCmd = &cobra.Command{
	Use:   "push [-i <modification-id> | --id <modification-id>] [--campaign-id <campaign-id>]",
	Short: "push modification code",
	Long:  `push modification code`,
	Run: func(cmd *cobra.Command, args []string) {
		var codeByte []byte
		var modification_ web_experimentation.Modification

		if !utils.CheckSingleFlag(filePath != "", code != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (file, code)")
		}

		campaignID, err := strconv.Atoi(CampaignID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		variationID_, err := strconv.Atoi(variationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		modificationID, err := strconv.Atoi(ModificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		modifList, err := httprequest.ModificationRequester.HTTPGetModification(campaignID, modificationID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		for _, modification := range modifList {
			if modification.Type == "customScriptNew" && modification.Selector != "" {
				modification_ = modification
			}
		}

		if filePath != "" {
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			codeByte = fileContent
		}

		if code != "" {
			codeByte = []byte(code)
		}

		selector_ := modification_.Selector
		if selector != "" {
			selector_ = selector
		}

		if modification_ == (web_experimentation.Modification{}) {
			if variationID_ == 0 && selector == "" {
				log.Fatalf("error occurred: Flag variation-id and selector are required.")
			}

			modificationToPush := web_experimentation.ModificationCodeCreateStruct{
				InputType:   "modification",
				Name:        "",
				Value:       string(codeByte),
				Selector:    selector,
				Type:        "customScriptNew",
				Engine:      "",
				VariationID: variationID_,
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
			Name:      modification_.Name,
			Value:     string(codeByte),
			Selector:  selector_,
			Type:      modification_.Type,
			Engine:    modification_.Engine,
		}

		body, err := httprequest.ModificationRequester.HTTPEditModification(campaignID, modification_.Id, modificationToPush)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(body))
	},
}

func init() {
	pushCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the campaign")

	if err := pushCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCmd.Flags().StringVarP(&ModificationID, "id", "i", "", "id of the modification code")

	if err := pushCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCmd.Flags().StringVarP(&variationID, "variation-id", "", "", "id of the variation")
	pushCmd.Flags().StringVarP(&code, "code", "c", "", "new code to push in the modification")
	pushCmd.Flags().StringVarP(&selector, "selector", "", "", "new selector to push in the modification")
	pushCmd.Flags().StringVarP(&filePath, "file", "", "", "file that contains new code to push in the modification")

	ModificationCodeCmd.AddCommand(pushCmd)
}
