/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpCreateCampaign(DataRaw)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data_raw", "", "", "the data")

	if err := createCmd.MarkFlagRequired("data_raw"); err != nil {
		fmt.Println(err)
	}

	CampaignCmd.AddCommand(createCmd)
}
