/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package targeting_key

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <targeting-key-id> | --id=<targeting-key-id>]",
	Short: "Delete a targeting key",
	Long:  `Delete a targeting key in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteTargetingKey(TargetingKeyID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Targeting key deleted")

	},
}

func init() {
	deleteCmd.Flags().StringVarP(&TargetingKeyID, "id", "i", "", "id of the targeting key you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
	TargetingKeyCmd.AddCommand(deleteCmd)
}
