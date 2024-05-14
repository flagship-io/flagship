/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account_global_code

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
	Use:   "push [-i <account-id> | --id <account-id>]",
	Short: "Push account global code",
	Long:  `push account global code`,
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

		body, err := httprequest.AccountGlobalCodeRequester.HTTPPushAccountGlobalCode(AccountID, codeByte)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(body))
	},
}

func init() {
	pushCmd.Flags().StringVarP(&AccountID, "id", "i", "", "id of the global code account")
	if err := pushCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	pushCmd.Flags().StringVarP(&code, "code", "c", "", "new code to push in the account")
	pushCmd.Flags().StringVarP(&filePath, "file", "", "", "file that contains new code to push in the account")

	AccountGlobalCodeCmd.AddCommand(pushCmd)
}
