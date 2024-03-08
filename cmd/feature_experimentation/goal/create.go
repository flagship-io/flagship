/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package goal

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Create a goal",
	Long:  `Create a goal in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPCreateGoal(DataRaw)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your goal, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	GoalCmd.AddCommand(createCmd)
}
