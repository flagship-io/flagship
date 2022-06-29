/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func check(token string) string {
	return "checking the token " + token
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "this authorization check the token",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetViper().GetString("token") != "" {
			fmt.Println(check(viper.GetViper().GetString("token")))
		} else {
			fmt.Println("token required")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(checkCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
