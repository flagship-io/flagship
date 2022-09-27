/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [--campaign-id=<campaign-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Create a variation group",
	Long:  `Create a variation group in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateVariationGroup(CampaignID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "variation group created: %s\n", body)
	},
}

func init() {
	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your variation group, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationGroupCmd.AddCommand(createCmd)
}
