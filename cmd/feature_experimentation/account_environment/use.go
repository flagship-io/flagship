/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package accountenvironment

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
	Short: "use a specific account envrionment id",
	Long:  `use a specific account envrionment id`,
	Run: func(cmd *cobra.Command, args []string) {
		if AccountEnvironmentID == "" {
			fmt.Fprintln(cmd.OutOrStderr(), "required flag account-id or account-environment-id")
			return
		}

		config.SetAccountEnvID(utils.FEATURE_EXPERIMENTATION, AccountEnvironmentID)

		fmt.Fprintln(cmd.OutOrStdout(), "Account Environment ID set to : "+AccountEnvironmentID)

	},
}

func init() {
	useCmd.Flags().StringVarP(&AccountEnvironmentID, "id", "i", "", "account env id of the credentials you want to manage")

	if err := useCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	AccountEnvironmentCmd.AddCommand(useCmd)
}
