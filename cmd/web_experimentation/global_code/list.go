/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package global_code

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
	Short: "List all global code",
	Long:  `List all global code`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.GlobalCodeRequester.HTTPListGlobalCode()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		utils.FormatItem([]string{"Id", "Code"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	GlobalCodeCmd.AddCommand(listCmd)
}
