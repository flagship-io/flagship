/*
Copyright © 2022 Flagship Team flagship@abtasty.com

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
	Use:     "variation-group [create|edit|get|list|delete]",
	Aliases: []string{"vg"},
	Short:   "Manage your variation groups",
	Long:    `Manage your variation groups`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationGroupCmd.PersistentFlags().StringVarP(&CampaignID, "campaign-id", "", "", "id of the campaign where you want to manage your variation group")

	if err := VariationGroupCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
