/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package feature_experimentation

import (
	"github.com/flagship-io/flagship/cmd/feature_experimentation/analyze"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/campaign"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/flag"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/goal"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/panic"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/project"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/resource"
	targetingkey "github.com/flagship-io/flagship/cmd/feature_experimentation/targeting_key"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/user"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/variation"
	variationgroup "github.com/flagship-io/flagship/cmd/feature_experimentation/variation_group"

	"github.com/spf13/cobra"
)

// FeatureExperimentationCmd represents the feature experimentation command
var FeatureExperimentationCmd = &cobra.Command{
	Use:     "feature-experimentation [project|campaign|flag|goal|targeting-key|variation-group|variation]",
	Aliases: []string{"feature-experimentation", "feature-exp", "fe", "feat-exp"},
	Short:   "Manage resources related to the feature experimentation product",
	Long:    `Manage resources related to the feature experimentation product in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommandPalettes() {
	FeatureExperimentationCmd.AddCommand(campaign.CampaignCmd)
	FeatureExperimentationCmd.AddCommand(project.ProjectCmd)
	FeatureExperimentationCmd.AddCommand(panic.PanicCmd)
	FeatureExperimentationCmd.AddCommand(user.UserCmd)
	FeatureExperimentationCmd.AddCommand(variationgroup.VariationGroupCmd)
	FeatureExperimentationCmd.AddCommand(variation.VariationCmd)
	FeatureExperimentationCmd.AddCommand(flag.FlagCmd)
	FeatureExperimentationCmd.AddCommand(goal.GoalCmd)
	FeatureExperimentationCmd.AddCommand(targetingkey.TargetingKeyCmd)
	FeatureExperimentationCmd.AddCommand(analyze.AnalyzeCmd)
	FeatureExperimentationCmd.AddCommand(resource.ResourceCmd)
}

func init() {
	addSubCommandPalettes()
}
