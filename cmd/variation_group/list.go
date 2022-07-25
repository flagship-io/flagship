/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ListVariationGroup() string {
	return "list variation groups of all campaign: \n campaign_id: " + viper.GetString("campaign_id") + "\n account_env_id: " + viper.GetString("account_environment_id")
}

// createCmd represents the create command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ListVariationGroup())
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	VariationGroupCmd.AddCommand(listCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
