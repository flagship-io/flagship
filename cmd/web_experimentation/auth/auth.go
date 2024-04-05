/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

import "github.com/spf13/cobra"

var (
	ClientID     string
	ClientSecret string
	Username     string
)

// AuthCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:     "authentication [login|get|list|delete]",
	Aliases: []string{"auth"},
	Short:   "Manage your CLI authentication for web experimentation",
	Long:    `Manage your CLI authentication for web experimentation`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
