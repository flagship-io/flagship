/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nameToggle string
)

func toggleCampaign(campaign string) string {
	return "toggle campaign " + campaign
}

// createCmd represents the create command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "this toggle campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(toggleCampaign(nameToggle))
	},
}

func init() {

	toggleCmd.Flags().StringVarP(&nameToggle, "name", "n", "", "the name")

	if err := toggleCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(toggleCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
