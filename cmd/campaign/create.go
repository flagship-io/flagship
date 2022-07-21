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

	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(createCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
