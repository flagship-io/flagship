/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account

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
	Short: "list all auth",
	Long:  `list all auth from your system`,
	Run: func(cmd *cobra.Command, args []string) {

		body, err := httprequest.CampaignRequester.HTTPListCampaign()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"Id", "ProjectId", "Name", "Description", "Type", "Status"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {

	AccountCmd.AddCommand(listCmd)
}
