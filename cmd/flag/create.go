/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [--name <flag-name> --type <flag-type> --default-value <default-value> --description <description> | -d <data-raw> | --data-raw <data-raw>]",
	Short: "Create a flag",
	Long:  `Create a flag in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		var data string
		var predefinedValues_ []string

		if DataRaw != "" && FlagName != "" {
			log.Fatalln("error occurred: You can either use --data-raw or the flags --name --description --type --default-value to create a flag")
		}

		if DataRaw != "" {
			data = DataRaw
		} else {
			if FlagPredefinedValues != "" {
				err := json.Unmarshal([]byte(FlagPredefinedValues), &predefinedValues_)
				if err != nil {
					log.Fatalf("error occurred1: %s", err)
				}
			}

			data_, err := json.Marshal(models.Flag{
				Name:             FlagName,
				Type:             FlagType,
				Description:      FlagDescription,
				DefaultValue:     FlagDefaultValue,
				Source:           "cli",
				PredefinedValues: predefinedValues_,
			})

			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			data = string(data_)
		}
		body, err := httprequest.HTTPCreateFlag(data)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", body)
	},
}

func init() {

	createCmd.Flags().StringVarP(&FlagName, "name", "", "", "name of the flag")
	createCmd.Flags().StringVarP(&FlagType, "type", "", "", "type of the flag")
	createCmd.Flags().StringVarP(&FlagDescription, "description", "", "Flag created from the CLI", "description of the flag")
	createCmd.Flags().StringVarP(&FlagDefaultValue, "default-value", "", "", "default value of the flag")
	createCmd.Flags().StringVarP(&FlagPredefinedValues, "predefined-values", "", "", "predefined valued for the flag")

	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your flag, check the doc for details")

	FlagCmd.AddCommand(createCmd)
}
