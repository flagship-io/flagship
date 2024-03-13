/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package variation

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	TestID      int
	VariationID int
)

// VariationCmd represents the variation command
var VariationCmd = &cobra.Command{
	Use:   "variation [create|edit|get|delete]",
	Short: "Manage your variation",
	Long:  `Manage your variation in your test`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	VariationCmd.PersistentFlags().IntVarP(&TestID, "test-id", "", 0, "id of the campaign where you want to manage your variation group")

	if err := VariationCmd.MarkPersistentFlagRequired("test-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
