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
	deleteVariationGroupId string
)

func deleteVariationGroup(variationGroupId string) string {
	return "delete variation group \n campaign_id: " + viper.GetString("campaign_id") + "\n variation_group_id: " + variationGroupId + "\n account_env_id: " + viper.GetString("account_environment_id")
}

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(deleteVariationGroup(deleteVariationGroupId))
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&deleteVariationGroupId, "variation_group_id", "i", "", "get variation group by id")

	if err := deleteCmd.MarkFlagRequired("variation_group_id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	VariationGroupCmd.AddCommand(deleteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
