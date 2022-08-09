/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package variation

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	CampaignID       string
	VariationGroupID string
	VariationID      string
	DataRaw          string
)

// VariationCmd represents the variation command
var VariationCmd = &cobra.Command{
	Use:   "variation [create|edit|get|list|delete]",
	Short: "Manage your variations",
	Long:  `Manage your variations in your variation group`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationCmd.PersistentFlags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the campaign where you want to manage your variation")

	if err := VariationCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationCmd.PersistentFlags().StringVarP(&VariationGroupID, "variation-group-id", "", "", "id of the variation group where you want to manage your variation")

	if err := VariationCmd.MarkPersistentFlagRequired("variation-group-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
}
