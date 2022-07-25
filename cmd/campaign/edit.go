/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditCampaign(CampaignID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("campaign updated: %s", body)
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
