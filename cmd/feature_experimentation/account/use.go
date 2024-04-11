/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account

import (
	"fmt"
	"log"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
)

// getCmd represents the list command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "get an auth credential",
	Long:  `list an auth credential`,
	Run: func(cmd *cobra.Command, args []string) {
		if AccountID == "" {
			fmt.Fprintln(cmd.OutOrStderr(), "required flag account-id")
			return
		}

		config.SetAccountID(utils.FEATURE_EXPERIMENTATION, AccountID)

		fmt.Fprintln(cmd.OutOrStdout(), "Account ID set to : "+AccountID)

	},
}

func init() {
	useCmd.Flags().StringVarP(&AccountID, "id", "i", "", "account id of the credentials you want to manage")

	if err := useCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	AccountCmd.AddCommand(useCmd)
}
