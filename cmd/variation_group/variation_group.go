/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package variation_group

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Account_environement_id string
	cfgFile_variation_group string
	campaign_id             string
)

// VariationGroupCmd represents the variation command
var VariationGroupCmd = &cobra.Command{
	Use:   "variation_group",
	Short: "variation group short desc",
	Long:  `variation group long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)
	VariationGroupCmd.PersistentFlags().StringVarP(&campaign_id, "campaign_id", "", "", "campaign_id")
	viper.BindPFlag("campaign_id", VariationGroupCmd.PersistentFlags().Lookup("campaign_id"))

	VariationGroupCmd.PersistentFlags().StringVarP(&cfgFile_variation_group, "data", "d", "", "config file (default is $PWD/variation_group.yaml)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLocalConfig() {

	if cfgFile_variation_group != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile_variation_group)
	} else {
		// Find home directory.
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship-mock" (without extension).
		viper.SetConfigFile(workingDir + "/variation_group.json")
	}
	// read in environment variables that match

	// If a config file is found, read it in.
	viper.MergeInConfig()
}
