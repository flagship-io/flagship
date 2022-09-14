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

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [--campaign-id=<campaign-id>] [-i <variation-group-id> | --id <variation-group-id>]",
	Short: "Get a variation group",
	Long:  `Get a variation group in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetVariationGroup(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	getCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "id of the variation group you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	VariationGroupCmd.AddCommand(getCmd)
}
