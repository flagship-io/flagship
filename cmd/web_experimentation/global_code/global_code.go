/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package global_code

import (
	"github.com/spf13/cobra"
)

// globalCodeCmd represents the global code command
var GlobalCodeCmd = &cobra.Command{
	Use:   "global-code [get|list]",
	Short: "Manage your tests",
	Long:  `Manage your tests`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
