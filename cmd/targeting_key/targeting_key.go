/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package targeting_key

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
	Long:    `Manage your targeting keys in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
