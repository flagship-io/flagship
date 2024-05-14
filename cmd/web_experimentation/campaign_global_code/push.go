/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign_global_code

import (
	"fmt"
	"log"
	"os"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

var code string
var filePath string

// pushCmd represents push command
var pushCmd = &cobra.Command{
	Use:   "push [-i <campaign-id> | --id <campaign-id>]",
	Short: "push campaign global code",
	Long:  `push campaign global code`,
	Run: func(cmd *cobra.Command, args []string) {
		var codeByte []byte

		if !utils.CheckSingleFlag(filePath != "", code != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (file, code)")
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

		body, err := httprequest.CampaignGlobalCodeRequester.HTTPPushCampaignGlobalCode(CampaignID, codeByte)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(body))
	},
}

func init() {
	pushCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the global code campaign")
	if err := pushCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCmd.Flags().StringVarP(&code, "code", "c", "", "new code to push in the campaign")
	pushCmd.Flags().StringVarP(&filePath, "file", "", "", "file that contains new code to push in the campaign")

	CampaignGlobalCodeCmd.AddCommand(pushCmd)
}
