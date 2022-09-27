/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [--campaign-id=<campaign-id>] [--variation-group-id=<variation-group-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Create a variation",
	Long:  `Create a variation in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateVariation(CampaignID, VariationGroupID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "variation created: %s\n", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your variation, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationCmd.AddCommand(createCmd)
}
