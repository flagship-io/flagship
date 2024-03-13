/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package campaign

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request/web_experimentation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all campaigns",
	Long:  `List all campaigns`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListTest()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "Name", "Description", "Type", "State"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	CampaignCmd.AddCommand(listCmd)
}