/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package goal

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all goals",
	Long:  `List all goals`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.GoalRequester.HTTPListGoal()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Label", "Type", "Operator", "Value"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	GoalCmd.AddCommand(listCmd)
}
