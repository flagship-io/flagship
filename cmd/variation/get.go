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

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get variation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetVariation(CampaignID, VariationGroupID, VariationID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Reference", "Allocation"}, body, viper.GetString("output_format"))

	},
}

func init() {
	getCmd.Flags().StringVarP(&VariationID, "id", "i", "", "get variation by id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	VariationCmd.AddCommand(getCmd)
}
