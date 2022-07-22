/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package panic

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var status string

func panicMessage(status string) string {
	return "panic the account env id: " + status
}

// campaignCmd represents the campaign command
var PanicCmd = &cobra.Command{
	Use:   "panic",
	Short: "panic short desc",
	Long:  `panic long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(panicMessage(viper.GetViper().GetString("account_environment_id")))
	},
}

func init() {

	PanicCmd.Flags().StringVarP(&status, "status", "s", "", "panic mode")
	if err := PanicCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
}
