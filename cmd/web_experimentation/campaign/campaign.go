/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign

import (
	"github.com/spf13/cobra"
)

var (
	CampaignID string
	Status     string
	DataRaw    string
)

// campaignCmd represents the campaign command
var CampaignCmd = &cobra.Command{
	Use:   "campaign [get|list|delete|switch]",
	Short: "Manage your campaigns",
	Long:  `Manage your campaigns`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
