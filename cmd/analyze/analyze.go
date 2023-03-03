/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package analyze

import (
	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/flagship/cmd/analyze/flag"
	"github.com/spf13/cobra"
)

var FSConfig *config.Config

// analyzeCmd represents the analyze command
var AnalyzeCmd = &cobra.Command{
	Use:   "analyze [flag]",
	Short: "Analyze your codebase",
	Long:  `Analyze your codebase using the codebase analyzer`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AnalyzeCmd.AddCommand(flag.FlagCmd)
}
