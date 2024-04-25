package utils

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecuteCommand(t *testing.T) {
	mockCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Hello, World!")
		},
	}

	output, err := ExecuteCommand(mockCmd)
	if err != nil {
		t.Errorf("ExecuteCommand() error = %v", err)
	}

	expectedOutput := "Hello, World!\n"
	if output != expectedOutput {
		t.Errorf("ExecuteCommand() returned unexpected output. Got: %s, Want: %s", output, expectedOutput)
	}
}
