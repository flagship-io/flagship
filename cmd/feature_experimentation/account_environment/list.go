/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package accountenvironment

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accountID string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all account environment id associated to your account",
	Long:  `list all account environment id associated to your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.AccountEnvironmentFERequester.HTTPListAccountEnvironment(accountID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Environment", "IsMain", "Panic", "SingleAssignment"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	listCmd.Flags().StringVarP(&accountID, "account-id", "a", "", "account id of the credentials you want to list")
	AccountEnvironmentCmd.AddCommand(listCmd)
}
