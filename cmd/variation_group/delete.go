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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteVariationGroup(CampaignID, VariationGroupID)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Variation group deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&VariationGroupID, "id", "i", "", "delete variation group by id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	VariationGroupCmd.AddCommand(deleteCmd)
}
