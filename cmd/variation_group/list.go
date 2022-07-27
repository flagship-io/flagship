/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListVariationGroup(CampaignID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"))
	},
}

func init() {

	listCmd.Flags().StringVarP(&CampaignID, "campaign-id", "", "", "campaign_id")

	if err := listCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationGroupCmd.AddCommand(listCmd)
}
