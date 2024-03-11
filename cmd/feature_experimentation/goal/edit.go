/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package goal

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request/feature_experimentation"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-i <goal-id> | --id=<goal-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Edit a goal",
	Long:  `Edit a goal in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditGoal(GoalID, DataRaw)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&GoalID, "id", "i", "", "id of the goal you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your goal, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	GoalCmd.AddCommand(editCmd)
}
