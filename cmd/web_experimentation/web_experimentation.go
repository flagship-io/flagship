/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package web_experimentation

import (
	"os"

	"github.com/flagship-io/flagship/cmd/web_experimentation/account"
	account_global_code "github.com/flagship-io/flagship/cmd/web_experimentation/account_global_code"
	"github.com/flagship-io/flagship/cmd/web_experimentation/auth"
	"github.com/flagship-io/flagship/cmd/web_experimentation/campaign"
	campaign_global_code "github.com/flagship-io/flagship/cmd/web_experimentation/campaign_global_code"
	"github.com/flagship-io/flagship/cmd/web_experimentation/variation"
	variation_global_code "github.com/flagship-io/flagship/cmd/web_experimentation/variation_global_code"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/flagship-io/flagship/utils/http_request"
	"github.com/flagship-io/flagship/utils/http_request/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// WebExperimentationCmd represents the web experimentation command
var WebExperimentationCmd = &cobra.Command{
	Use:     "web-experimentation [auth|account|campaign|global-code|variation]",
	Aliases: []string{"web-experimentation", "web-exp", "we"},
	Short:   "Manage resources related to the feature experimentation product",
	Long:    `Manage resources related to the feature experimentation product`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommandPalettes() {
	WebExperimentationCmd.AddCommand(campaign.CampaignCmd)
	WebExperimentationCmd.AddCommand(variation.VariationCmd)
	WebExperimentationCmd.AddCommand(auth.AuthCmd)
	WebExperimentationCmd.AddCommand(account.AccountCmd)
	WebExperimentationCmd.AddCommand(campaign_global_code.CampaignGlobalCodeCmd)
	WebExperimentationCmd.AddCommand(account_global_code.AccountGlobalCodeCmd)
	WebExperimentationCmd.AddCommand(variation_global_code.VariationGlobalCodeCmd)
}

func init() {
	addSubCommandPalettes()
}

func initConfig() {
	v := viper.New()
	homeDir, _ := os.UserHomeDir()
	var requestConfig = common.RequestConfig{Product: utils.WEB_EXPERIMENTATION}

	v.SetConfigFile(homeDir + "/.flagship/credentials/" + utils.WEB_EXPERIMENTATION + "/.cli.yaml")
	v.MergeInConfig()
	if v.GetString("current_used_credential") != "" {
		vL := config.ReadAuth(utils.WEB_EXPERIMENTATION, v.GetString("current_used_credential"))
		v.MergeConfigMap(vL.AllSettings())
	}

	v.Unmarshal(&requestConfig)
	common.Init(requestConfig)

	r := &http_request.ResourceRequester

	r.Init(&requestConfig)
	return
}
