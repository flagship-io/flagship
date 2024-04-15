/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var modifFiles ModificationGlobalCode

// generate represents generate command
var generateCmd = &cobra.Command{
	Use:   "generate-file [-i <variation-id> | --id <variation-id>] [--campaign-id <campaign-id>]",
	Short: "Generate variation global code file",
	Long:  `Generate variation global code file`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := GetCodeFiles(VariationID, CampaignID)

		if resp.JS != "" {
			jsFilePath := strconv.Itoa(VariationID) + ".js"

			err := os.WriteFile(jsFilePath, []byte(resp.JS), os.ModePerm)
			if err != nil {
				fmt.Println("Error writing JavaScript file:", err)
				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), "JavaScript code file generated successfully:", jsFilePath)
		}

		if resp.CSS != "" {
			cssFilePath := strconv.Itoa(VariationID) + ".css"

			err := os.WriteFile(cssFilePath, []byte(resp.CSS), os.ModePerm)
			if err != nil {
				fmt.Println("Error writing CSS file:", err)
				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), "CSS code file generated successfully:", cssFilePath)
		}

	},
}

func init() {
	generateCmd.Flags().IntVarP(&CampaignID, "campaign-id", "", 0, "id of the global code campaign you want to display")

	if err := generateCmd.MarkFlagRequired("campaign-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	generateCmd.Flags().IntVarP(&VariationID, "id", "i", 0, "id of the global code vairation you want to display")

	if err := generateCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	VariationGlobalCodeCmd.AddCommand(generateCmd)
}
