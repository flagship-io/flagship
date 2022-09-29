/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Use:   "campaign [create|edit|get|list|delete|toggle]",
	Short: "Manage your campaigns",
	Long:  `Manage your campaigns`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
