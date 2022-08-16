/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteCampaign(CampaignID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Campaign deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "delete the campaign")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(deleteCmd)
}
