/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign_global_code

import (
	"github.com/spf13/cobra"
)

// CampaignGlobalCodeCmd represents the campaign global code command
var CampaignGlobalCodeCmd = &cobra.Command{
	Use:     "campaign-global-code [get]",
	Short:   "Get campaign global code",
	Aliases: []string{"cgc"},
	Long:    `Get campaign global code from your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
