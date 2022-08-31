/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package targeting_key

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-i <targeting-key-id> | --id=<targeting-key-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Edit a targeting key",
	Long:  `Edit a targeting key in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditTargetingKey(TargetingKeyID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("Targeting key updated: %s", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&TargetingKeyID, "id", "i", "", "id of the targeting key you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your targeting key, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	TargetingKeyCmd.AddCommand(editCmd)
}
