/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account_global_code

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

var accountID string

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <account-id> | --id <account-id>]",
	Short: "Get global account code",
	Long:  `Get global account code`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.AccountGlobalCodeRequester.HTTPGetAccountGlobalCode(accountID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), body)

	},
}

func init() {
	getCmd.Flags().StringVarP(&accountID, "id", "i", "", "id of the global code account you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	AccountGlobalCodeCmd.AddCommand(getCmd)
}
