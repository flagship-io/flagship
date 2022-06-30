/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deleteCampaignId string
)

func deleteCampaign(campaign_id string) string {
	return "delete campaign \n campaign_id: " + campaign_id
}

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(deleteCampaign(deleteCampaignId))
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&deleteCampaignId, "campaign_id", "i", "", "delete the campaign")

	if err := deleteCmd.MarkFlagRequired("campaign_id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(deleteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
