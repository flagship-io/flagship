/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "this toggle campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !(Status == "active" || Status == "paused" || Status == "interrupted") {
			fmt.Println("Status can only have 3 values : active or paused or interrupted")
		} else {
			httprequest.HttpToggleCampaign(CampaignID, Status)
		}

	},
}

func init() {

	toggleCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "toggle campaign id")
	toggleCmd.Flags().StringVarP(&Status, "status", "s", "", "status")

	if err := toggleCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	if err := toggleCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}

	CampaignCmd.AddCommand(toggleCmd)
}
