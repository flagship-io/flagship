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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete variation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteVariation(CampaignID, VariationGroupID, VariationID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("variation deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&VariationID, "id", "i", "", "delete variation by id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	VariationCmd.AddCommand(deleteCmd)
}
