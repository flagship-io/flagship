/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will match the tag
var Version = "v0.1.0"

func DisplayVersion() {
	fmt.Printf("Flagship CLI verison : %s\n", Version)
}

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version short desc",
	Long:  `version long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		DisplayVersion()
	},
}
