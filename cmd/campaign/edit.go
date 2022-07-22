/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpEditCampaign(CampaignID, DataRaw)
	},
}

func init() {

	editCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "edit the campaign")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "the data")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		fmt.Println(err)
	}

	CampaignCmd.AddCommand(editCmd)
}
