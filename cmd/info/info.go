/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package info

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InfoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the information like account environment id, account id and client id",
	Long:  `Show the information like account environment id, account id and client id your Flagship CLI runs on`,
	Run: func(cmd *cobra.Command, args []string) {

		yellow := color.New(color.FgYellow).SprintFunc()

		fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", yellow("Account ID"), viper.GetString("account_id"))
		fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", yellow("Account environment ID"), viper.GetString("account_environment_id"))
		fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", yellow("Client ID"), viper.GetString("client_id"))
		fmt.Fprintf(cmd.OutOrStdout(), "%s: %d (in second)\n", yellow("Token expiration"), viper.GetInt("expiration"))
		fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", yellow("Token scope"), viper.GetString("scope"))
	},
}
