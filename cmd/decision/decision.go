/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package decision

import (
	"github.com/flagship-io/flagship/cmd/decision/self_hosted"
	"github.com/spf13/cobra"
)

// DecisionCmd represents the decision command
var DecisionCmd = &cobra.Command{
	Use:   "decision",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	DecisionCmd.AddCommand(self_hosted.DecisionSelfHostedCmd)
}
