/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package right

import (
	"github.com/spf13/cobra"
)

// campaignCmd represents the campaign command
var RightCmd = &cobra.Command{
	Use:   "right",
	Short: "right short desc",
	Long:  `right long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
