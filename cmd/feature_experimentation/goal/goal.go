/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package goal

import (
	"github.com/spf13/cobra"
)

var (
	GoalID  string
	DataRaw string
)

// GoalCmd represents the goal command
var GoalCmd = &cobra.Command{
	Use:   "goal [create|edit|get|list|delete]",
	Short: "Manage your goals",
	Long:  `Manage your goals`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
