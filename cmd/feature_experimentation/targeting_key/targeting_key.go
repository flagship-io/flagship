/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package targetingkey

import "github.com/spf13/cobra"

var (
	TargetingKeyID string
	DataRaw        string
)

// TargetingKeyCmd represents the targeting command
var TargetingKeyCmd = &cobra.Command{
	Use:     "targeting-key [create|edit|get|list|delete]",
	Aliases: []string{"tk"},
	Short:   "Manage your targeting keys",
	Long:    `Manage your targeting keys`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
