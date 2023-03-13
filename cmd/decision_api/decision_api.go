/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package decision_api

import (
	"log"

	"github.com/flagship-io/flagship/cmd/decision_api/cloud_hosted"
	"github.com/flagship-io/flagship/cmd/decision_api/self_hosted"
	"github.com/spf13/cobra"
)

var (
	VisitorId      string
	VisitorContext []string
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

	DecisionCmd.Flags().StringVarP(&VisitorId, "visitor-id", "", "", "visitorId")

	if err := DecisionCmd.MarkFlagRequired("visitor-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	DecisionCmd.AddCommand(self_hosted.DecisionSelfHostedCmd)
	DecisionCmd.AddCommand(cloud_hosted.DecisionCloudHostedCmd)
}
