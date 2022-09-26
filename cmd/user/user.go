/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package user

import (
	"github.com/spf13/cobra"
)

var (
	UserEmail string
	UserRole  string
	DataRaw   string
)

// userCmd represents the user command
var UserCmd = &cobra.Command{
	Use:   "user [create|edit|list|delete]",
	Short: "Manage your users",
	Long:  `Manage your users`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
