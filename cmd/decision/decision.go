/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package decision

import (
	"github.com/spf13/cobra"
)

// DecisionCmd represents the decision command
var DecisionCmd = &cobra.Command{
	Use:   "decision [flag]",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
