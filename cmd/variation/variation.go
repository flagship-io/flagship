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
	Use:   "variation",
	Short: "variation short desc",
	Long:  `variation long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationCmd.PersistentFlags().StringVarP(&CampaignID, "campaign-id", "", "", "campaign_id")

	if err := VariationCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	VariationCmd.PersistentFlags().StringVarP(&VariationGroupID, "variation-group-id", "", "", "variation_group_id")

	if err := VariationCmd.MarkPersistentFlagRequired("variation-group-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
}
