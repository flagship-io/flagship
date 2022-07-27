/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package variation_group

import (
	"github.com/spf13/cobra"
)

var (
	CampaignID       string
	VariationGroupID string
)

// VariationGroupCmd represents the variation command
var VariationGroupCmd = &cobra.Command{
	Use:     "variationgroup",
	Short:   "variation group short desc",
	Aliases: []string{"vg"},
	Long:    `variation group long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
