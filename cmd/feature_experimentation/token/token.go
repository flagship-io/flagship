/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package token

import (
	"github.com/spf13/cobra"
)

// TokenCmd represents the token command
var TokenCmd = &cobra.Command{
	Use:   "token [info]",
	Short: "Manage your token",
	Long:  `Manage your token`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
