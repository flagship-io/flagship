package utils

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type HTTPListResponse[T any] struct {
	Items             []T `json:"items"`
	CurrentItemsCount int `json:"current_items_count"`
	CurrentPage       int `json:"current_page"`
	TotalCount        int `json:"total_count"`
	ItemsPerPage      int `json:"items_per_page"`
	LastPage          int `json:"last_page"`
}

func ExecuteCommand(cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err = cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %s", err)

	}

	return buf.String(), err
}
