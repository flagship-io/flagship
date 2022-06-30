/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getCampaignId string
)

func getCampaign(campaign_id string) string {
	return "get campaign \n campaign_id: " + campaign_id
}

// createCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getCampaign(getCampaignId))
	},
}

func init() {

	getCmd.Flags().StringVarP(&getCampaignId, "campaign_id", "i", "", "edit campaign by campaign_id")

	if err := getCmd.MarkFlagRequired("campaign_id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(getCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
