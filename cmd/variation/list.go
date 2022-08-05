/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--campaign-id=<campaign-id>] [--variation-group-id=<variation-group-id>]",
	Short: "List a variation",
	Long:  `List a variation in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListVariation(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Reference", "Allocation"}, body, viper.GetString("output_format"))
	},
}

func init() {
	VariationCmd.AddCommand(listCmd)
}
