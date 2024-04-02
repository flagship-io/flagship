/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account

import "github.com/spf13/cobra"

var (
	Username             string
	AccountID            string
	AccountEnvironmentID string
)

// ConfigurationCmd represents the configuration command
var AccountCmd = &cobra.Command{
	Use:   "account [use|list|current]",
	Short: "Manage your CLI authentication",
	Long:  `Manage your CLI authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
