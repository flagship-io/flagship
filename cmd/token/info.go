/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package token

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get the information related to your token",
	Long:  `Get the information related to your token`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("token") != "" {
			body, err := httprequest.HTTPCheckToken(viper.GetString("token"))
			if err != nil {
				log.Fatalf("error occured: %v", err)
			}
			utils.FormatItem([]string{"ClientID", "AccountID", "ExpiresIn", "Scope"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
		} else {
			log.Println("Token required")
		}
	},
}

func init() {
	TokenCmd.AddCommand(infoCmd)
}
