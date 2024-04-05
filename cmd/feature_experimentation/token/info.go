/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package token

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get the information related to your token",
	Long:  `Get the information related to your token`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := common.HTTPCheckToken()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"ClientID", "AccountID", "ExpiresIn", "Scope"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	TokenCmd.AddCommand(infoCmd)
}
