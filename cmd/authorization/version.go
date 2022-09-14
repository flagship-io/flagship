/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package authorization

import (
	"io"
	"log"

	"github.com/spf13/cobra"
)

// Version will match the tag
var Version = "v0.2.1"

type Commad struct {
	*cobra.Command
}

func CmdWriter(w io.Writer) error {
	i := "Flagship CLI version : " + Version
	_, err := w.Write([]byte(i))
	return err
}

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version short desc",
	Long:  `version long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		err := CmdWriter(cmd.OutOrStdout())
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}
