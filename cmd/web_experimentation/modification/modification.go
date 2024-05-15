/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package modification

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	CampaignID     int
	ModificationID int
	Status         string
	DataRaw        string
)

// modificationCmd represents the modification command
var ModificationCmd = &cobra.Command{
	Use:   "modification [get|list|delete]",
	Short: "Manage your modifications",
	Long:  `Manage your modifications`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ModificationCmd.PersistentFlags().IntVarP(&CampaignID, "campaign-id", "", 0, "id of the campaign where you want to manage your modifications")

	if err := ModificationCmd.MarkPersistentFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
