/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	toggleCampaignId, statusCampaign string
)

func toggleCampaign(toggleCampaignId, statusCampaign string) string {
	return "toggle campaign \n campaign_id: " + toggleCampaignId + "\n status: " + statusCampaign
}

// createCmd represents the create command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "this toggle campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !(statusCampaign == "active" || statusCampaign == "paused" || statusCampaign == "interrupted") {
			fmt.Println("Status can only have 3 values : active or paused or interrupted")
		} else {
			fmt.Println(toggleCampaign(toggleCampaignId, statusCampaign))
		}

	},
}

func init() {

	toggleCmd.Flags().StringVarP(&toggleCampaignId, "campaign_id", "i", "", "toggle campaign id")
	toggleCmd.Flags().StringVarP(&statusCampaign, "status", "s", "", "status")

	if err := toggleCmd.MarkFlagRequired("campaign_id"); err != nil {
		fmt.Println(err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(toggleCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
