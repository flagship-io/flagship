/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation

import (
	"encoding/json"
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [--test-id=<test-id>] [-i=<variation-id> | --id=<variation-id>]",
	Short: "Get a variation",
	Long:  `Get a variation in your campaign`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.VariationWERequester.HTTPGetVariation(TestID, VariationID)
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
	getCmd.Flags().IntVarP(&VariationID, "id", "i", 0, "id of the variation group you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	VariationCmd.AddCommand(getCmd)
}
