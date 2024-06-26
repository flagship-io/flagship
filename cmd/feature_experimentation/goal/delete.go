/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package goal

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <goal-id> | --id=<goal-id>]",
	Short: "Delete a flag",
	Long:  `Delete a flag`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.GoalRequester.HTTPDeleteGoal(GoalID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Goal deleted")

	},
}

func init() {
	deleteCmd.Flags().StringVarP(&GoalID, "id", "i", "", "id of the goal you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	GoalCmd.AddCommand(deleteCmd)
}
