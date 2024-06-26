/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package feature_experimentation

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/cmd/feature_experimentation/account"
	accountenvironment "github.com/flagship-io/flagship/cmd/feature_experimentation/account_environment"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/analyze"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/auth"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/campaign"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/flag"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/goal"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/panic"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/project"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/resource"
	targetingkey "github.com/flagship-io/flagship/cmd/feature_experimentation/targeting_key"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/token"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/user"
	"github.com/flagship-io/flagship/cmd/feature_experimentation/variation"
	variationgroup "github.com/flagship-io/flagship/cmd/feature_experimentation/variation_group"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/flagship-io/flagship/utils/http_request"
	"github.com/flagship-io/flagship/utils/http_request/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FeatureExperimentationCmd represents the feature experimentation command
var FeatureExperimentationCmd = &cobra.Command{
	Use:     "feature-experimentation [auth|account|account-environment|project|campaign|flag|goal|targeting-key|variation-group|variation]",
	Aliases: []string{"feature-experimentation", "feature-exp", "fe", "feat-exp"},
	Short:   "Manage resources related to the feature experimentation product",
	Long:    `Manage resources related to the feature experimentation product in your account`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
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
	FeatureExperimentationCmd.AddCommand(auth.AuthCmd)
	FeatureExperimentationCmd.AddCommand(account.AccountCmd)
	FeatureExperimentationCmd.AddCommand(token.TokenCmd)
	FeatureExperimentationCmd.AddCommand(accountenvironment.AccountEnvironmentCmd)
}

func init() {
	addSubCommandPalettes()
}

func initConfig() {
	v := viper.New()
	homeDir, _ := os.UserHomeDir()
	var requestConfig = common.RequestConfig{Product: utils.FEATURE_EXPERIMENTATION}

	v.SetConfigFile(homeDir + "/.flagship/credentials/" + utils.FEATURE_EXPERIMENTATION + "/.cli.yaml")
	v.MergeInConfig()
	if v.GetString("current_used_credential") != "" {
		vL, err := config.ReadAuth(utils.FEATURE_EXPERIMENTATION, v.GetString("current_used_credential"))
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		v.MergeConfigMap(vL.AllSettings())
	}

	v.Unmarshal(&requestConfig)
	common.Init(requestConfig)
	resource.Init(requestConfig)
	viper.MergeConfigMap(v.AllSettings())

	r := &http_request.ResourceRequester

	r.Init(&requestConfig)
	return

}
