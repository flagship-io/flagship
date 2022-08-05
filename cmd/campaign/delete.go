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
	Use:   "delete [-i <campaign-id> | --campaign-id=<campaign-id>]",
	Short: "Delete a campaign",
	Long:  `Delete a campaign of your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteCampaign(CampaignID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("campaign deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "id of the campaign you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	CampaignCmd.AddCommand(deleteCmd)
}
