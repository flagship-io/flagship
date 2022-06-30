/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Account_environment_id string
	cfgFile_campaign       string
)

var (
	Name          string
	Project_id    string
	Description   string
	Campaign_type string
)

// campaignCmd represents the campaign command
var CampaignCmd = &cobra.Command{
	Use:   "campaign",
	Short: "campaign short desc",
	Long:  `campaign long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)
	CampaignCmd.PersistentFlags().StringVarP(&Account_environment_id, "account_environment_id", "a", "", "account environment id")
	viper.BindPFlag("account_environment_id", CampaignCmd.PersistentFlags().Lookup("account_environment_id"))

	CampaignCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "the campaign name")
	CampaignCmd.PersistentFlags().StringVarP(&Project_id, "project_id", "p", "", "the projact id")
	CampaignCmd.PersistentFlags().StringVarP(&Description, "description", "d", "", "the campaign description")
	CampaignCmd.PersistentFlags().StringVarP(&Campaign_type, "type", "t", "", "the campaign type")

	viper.BindPFlag("name", CampaignCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("project_id", CampaignCmd.PersistentFlags().Lookup("project_id"))
	viper.BindPFlag("description", CampaignCmd.PersistentFlags().Lookup("description"))
	viper.BindPFlag("type", CampaignCmd.PersistentFlags().Lookup("type"))

	CampaignCmd.PersistentFlags().StringVarP(&cfgFile_campaign, "config_campaign", "", "", "config file (default is $PWD/campaign.yaml)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLocalConfig() {

	if cfgFile_campaign != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile_campaign)
	} else {
		// Find home directory.
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship-mock" (without extension).
		viper.SetConfigFile(workingDir + "/campaign.json")
	}
	// read in environment variables that match

	// If a config file is found, read it in.
	viper.MergeInConfig()
}
