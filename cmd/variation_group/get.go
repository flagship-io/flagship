/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package variation_group

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	variationGroupId string
)

func getVariationGroup(variationGroupId string) string {
	return "get variation group \n campaign_id: " + viper.GetString("campaign_id") + "\n variation_group_id: " + variationGroupId + "\n account_env_id: " + viper.GetString("account_environment_id")
}

// createCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVariationGroup(variationGroupId))
	},
}

func init() {

	getCmd.Flags().StringVarP(&variationGroupId, "variation_group_id", "i", "", "get variation group by id")

	if err := getCmd.MarkFlagRequired("variation_group_id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	VariationGroupCmd.AddCommand(getCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
