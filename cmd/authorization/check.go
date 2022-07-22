/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "this authorization check the token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetViper().GetString("token") != "" {
			httprequest.HttpCheckToken(viper.GetViper().GetString("token"))
		} else {
			fmt.Println("token required")
		}
	},
}

func init() {
	AuthorizationCmd.AddCommand(checkCmd)
}
