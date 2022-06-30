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
	editvariationGroupId string
)

func editVariationGroup(variationGroupId string) string {
	return "edit variation group \n campaign_id: " + viper.GetViper().GetString("campaign_id") + "\n variation_group_id: " + variationGroupId + "\n new variation group from the data " + viper.GetViper().GetString("variation_groups") + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(editVariationGroup(editvariationGroupId))
	},
}

func init() {

	editCmd.Flags().StringVarP(&editvariationGroupId, "variation_group_id", "i", "", "edit variation group by id")

	if err := editCmd.MarkFlagRequired("variation_group_id"); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.
	VariationGroupCmd.AddCommand(editCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
