/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package panic

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Account_environment_id string
)

func panicMessage(account_environment_id string) string {
	return "panic the account env id: " + account_environment_id
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

	cobra.OnInitialize(initLocalConfig)

	PanicCmd.PersistentFlags().StringVarP(&Account_environment_id, "account_environment_id", "a", "", "account environment id")
	viper.BindPFlag("account_environment_id", PanicCmd.PersistentFlags().Lookup("account_environment_id"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLocalConfig() {
	viper.MergeInConfig()
}
