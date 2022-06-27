/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	editName string
)

func editCampaign(campaign string) string {
	return "edit campaign " + campaign
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(editCampaign(editName))
	},
}

func init() {

	editCmd.Flags().StringVarP(&editName, "name", "n", "", "the url to path")

	if err := editCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(editCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
