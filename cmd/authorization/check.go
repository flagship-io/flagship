/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "this authorization check the token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("token") != "" {
			err := httprequest.HTTPCheckToken(viper.GetString("token"))
			if err != nil {
				log.Fatalf("error occured: %v", err)
			}
			log.Println("Token ok")
		} else {
			log.Println("Token required")
		}
	},
}

func init() {
	AuthorizationCmd.AddCommand(checkCmd)
}
