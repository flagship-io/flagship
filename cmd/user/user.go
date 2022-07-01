/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"github.com/Chadiii/flagship-mock/cmd/user/right"
	"github.com/spf13/cobra"
)

// campaignCmd represents the campaign command
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "user short desc",
	Long:  `user long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	UserCmd.AddCommand(right.RightCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
