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
	Use:   "test [create|edit|get|list|delete|switch]",
	Short: "Manage your tests",
	Long:  `Manage your tests`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
