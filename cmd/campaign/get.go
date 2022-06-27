/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getName string
)

func getCampaign(campaign string) string {
	return "get campaign " + campaign
}

// createCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getCampaign(getName))
	},
}

func init() {

	getCmd.Flags().StringVarP(&getName, "name", "n", "", "the url to path")

	if err := getCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(getCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
