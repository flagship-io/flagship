/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpGetCampaign(CampaignID)
	},
}

func init() {

	getCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "get campaign by campaign_id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	CampaignCmd.AddCommand(getCmd)
}
