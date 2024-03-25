/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

import "github.com/spf13/cobra"

var (
	Username     string
	ClientID     string
	ClientSecret string
	AccountId    string
)

// ConfigurationCmd represents the configuration command
var AuthCmd = &cobra.Command{
	Use:   "auth [create|edit|get|list|delete|use]",
	Short: "Manage your CLI authentication",
	Long:  `Manage your CLI authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
