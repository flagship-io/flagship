package utils

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type HTTPListResponseFE[T any] struct {
	Items             []T `json:"items"`
	CurrentItemsCount int `json:"current_items_count"`
	CurrentPage       int `json:"current_page"`
	TotalCount        int `json:"total_count"`
	ItemsPerPage      int `json:"items_per_page"`
	LastPage          int `json:"last_page"`
}

type HTTPListResponseWE[T any] struct {
	Data       []T        `json:"_data"`
	Pagination Pagination `json:"_pagination"`
}

type Pagination struct {
	Total      int `json:"_total"`
	Pages      int `json:"_pages"`
	Page       int `json:"_page"`
	MaxPerPage int `json:"_max_per_page"`
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
