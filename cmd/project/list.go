/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Chadiii/flagship/models"
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

var outputFormat string

func formatTable(items models.FormatableItem) {
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)
	items.FormatTable(w)
	w.Flush()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HttpListProject()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		projects, err := httprequest.HttpListProjectFormat(body)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if outputFormat == "json" {
			projectJSON, err := json.Marshal(projects)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			fmt.Println(string(projectJSON))
			return
		}

		formatTable(models.ProjectItems(projects))
	},
}

func init() {

	listCmd.Flags().StringVarP(&outputFormat, "output-format", "f", "table", "the output format")

	ProjectCmd.AddCommand(listCmd)
}
