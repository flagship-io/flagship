package utils

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

func ExecuteCommand(cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err = cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}

	return buf.String(), err
}
