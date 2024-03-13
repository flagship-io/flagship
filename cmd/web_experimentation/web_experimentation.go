/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package web_experimentation

import (
	"github.com/flagship-io/flagship/cmd/web_experimentation/campaign"
	"github.com/flagship-io/flagship/cmd/web_experimentation/global_code"
	"github.com/flagship-io/flagship/cmd/web_experimentation/variation"

	"github.com/spf13/cobra"
)

// WebExperimentationCmd represents the web experimentation command
var WebExperimentationCmd = &cobra.Command{
	Use:     "web-experimentation [campaign]",
	Aliases: []string{"web-experimentation", "web-exp", "we"},
	Short:   "Manage resources related to the feature experimentation product",
	Long:    `Manage resources related to the feature experimentation product in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommandPalettes() {
	WebExperimentationCmd.AddCommand(campaign.CampaignCmd)
	WebExperimentationCmd.AddCommand(global_code.GlobalCodeCmd)
	WebExperimentationCmd.AddCommand(variation.VariationCmd)
}

func init() {
	addSubCommandPalettes()
}
