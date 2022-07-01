/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ListUser() string {
	return "list users: \n campaign_id: " + viper.GetViper().GetString("campaign_id") + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ListUser())
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	UserCmd.AddCommand(listCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
