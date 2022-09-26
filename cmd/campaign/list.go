/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package campaign

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all campaigns",
	Long:  `List all campaigns`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListCampaign()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "ProjectID", "Name", "Description", "Type", "Status"}, body, viper.GetString("output_format"))
	},
}

func init() {
	CampaignCmd.AddCommand(listCmd)
}
