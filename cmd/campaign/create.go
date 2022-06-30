/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func createCampaign() string {
	return "create campaign \n name: " + viper.GetViper().GetString("name") + "\n project_id: " + viper.GetViper().GetString("project_id") + "\n description: " + viper.GetViper().GetString("description") + "\n type: " + viper.GetViper().GetString("type") + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(createCampaign())
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(createCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
