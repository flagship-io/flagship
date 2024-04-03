/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account

import (
	"fmt"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
)

// getCmd represents the list command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "get an auth credential",
	Long:  `list an auth credential from your system`,
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			fmt.Fprintln(cmd.OutOrStdout(), "required flag account-id")
			return
		}

		if AccountID != "" {
			config.SetAccountID(utils.WEB_EXPERIMENTATION, AccountID)
		}

	},
}

func init() {
	useCmd.Flags().StringVarP(&AccountID, "account-id", "a", "", "account id of the credentials you want to display")

	AccountCmd.AddCommand(useCmd)
}
