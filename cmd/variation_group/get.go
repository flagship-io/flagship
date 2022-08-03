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
	Use:   "get",
	Short: "this get variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetVariationGroup(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"))

	},
}

func init() {
	getCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "get variation group by id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Println(err)
	}
	VariationGroupCmd.AddCommand(getCmd)
}
