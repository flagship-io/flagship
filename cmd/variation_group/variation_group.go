/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package variation_group

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	CampaignID       string
	VariationGroupID string
	DataRaw          string
)

// VariationGroupCmd represents the variation command
var VariationGroupCmd = &cobra.Command{
	Use:     "variationgroup [create|edit|get|list|delete]",
	Short:   "Manage your variation groups",
	Long:    `Manage your variation groups in your campaign`,
	Aliases: []string{"vg"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationGroupCmd.PersistentFlags().StringVarP(&CampaignID, "campaign-id", "", "", "campaign_id")

	if err := VariationGroupCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}
}
