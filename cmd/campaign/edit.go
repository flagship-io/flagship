/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	editCampaignId string
)

func editCampaign(campaign_id string) string {
	return "edit campaign \n campaign_id: " + campaign_id + "\n name: " + viper.GetViper().GetString("name") + "\n project_id: " + viper.GetViper().GetString("project_id") + "\n description: " + viper.GetViper().GetString("description") + "\n type: " + viper.GetViper().GetString("type") + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(editCampaign(editCampaignId))
	},
}

func init() {

	editCmd.Flags().StringVarP(&editCampaignId, "campaign_id", "i", "", "edit the campaign")

	if err := editCmd.MarkFlagRequired("campaign_id"); err != nil {
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
