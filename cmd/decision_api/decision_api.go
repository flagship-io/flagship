/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package decision_api

import (
	"github.com/flagship-io/flagship/cmd/decision_api/cloud_hosted"
	"github.com/flagship-io/flagship/cmd/decision_api/self_hosted"
	"github.com/spf13/cobra"
)

// DecisionCmd represents the decision command
var DecisionCmd = &cobra.Command{
	Use:   "decision-api",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	DecisionCmd.AddCommand(self_hosted.DecisionSelfHostedCmd)
	DecisionCmd.AddCommand(cloud_hosted.DecisionCloudHostedCmd)
}
