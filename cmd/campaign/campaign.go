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
	cfgFile                string
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

	cobra.OnInitialize(initConfig)

	CampaignCmd.PersistentFlags().StringVarP(&Account_environment_id, "account_environment_id", "a", "", "account environment id")
	viper.BindPFlag("account_environment_id", CampaignCmd.PersistentFlags().Lookup("account_environment_id"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship-mock" (without extension).
		viper.SetConfigFile(workingDir + "/campaign.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if cfgFile != "" {

		}
	}
}
