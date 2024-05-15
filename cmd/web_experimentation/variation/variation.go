/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package variation

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	CampaignID  int
	VariationID int
)

// VariationCmd represents the variation command
var VariationCmd = &cobra.Command{
	Use:   "variation [get|delete]",
	Short: "Manage your variation",
	Long:  `Manage your variation`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationCmd.PersistentFlags().IntVarP(&CampaignID, "campaign-id", "", 0, "id of the campaign where you want to manage your variation")

	if err := VariationCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
