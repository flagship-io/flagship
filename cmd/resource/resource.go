/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package resource

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	resourceFile string
)

var v = viper.New()

// ResourceCmd represents the resource command
var ResourceCmd = &cobra.Command{
	Use:   "resource [load]",
	Short: "Manage your resources",
	Long:  `Manage your resources`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
