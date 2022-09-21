/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package targeting_key

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Create a targeting key",
	Long:  `Create a targeting key in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateTargetingKey(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("targeting key created: %s", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your targeting key, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	TargetingKeyCmd.AddCommand(createCmd)
}
