/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package global_code

import (
	"encoding/json"
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request/web_experimentation"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all global code",
	Long:  `List all global code`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListGlobalCode()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		fmt.Printf("%s", string(jsonBody))

	},
}

func init() {
	GlobalCodeCmd.AddCommand(listCmd)
}
