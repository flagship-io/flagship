/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kyokomi/emoji/v2"
	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func summaryTableFlagCreated(flagCreatedLen, flagNotCreatedLen, flagAlreadyExistLen int) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	totalFlag := flagCreatedLen + flagAlreadyExistLen + flagNotCreatedLen

	summtbl := table.New("\nSummary")
	summtbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	summtbl.AddRow(fmt.Sprintf("Total flags: %d (%d Flag created %s, %d Flag not created%s, %d Flag that already exist%s)", totalFlag, flagCreatedLen, emoji.Sprint(":check_mark_button:"), flagNotCreatedLen, emoji.Sprint(":cross_mark:"), flagAlreadyExistLen, emoji.Sprint(":white_large_square:")))
	summtbl.Print()
}

func flagCreatedTable(cmd *cobra.Command, listedFlags []models.Flag) error {
	var flagCreatedLen int = 0
	var flagNotCreatedLen int = 0
	var flagAlreadyExistLen int = 0

	var flagKeyNotCreated []string

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Flag", "Type", "defaultValue", "File", fmt.Sprintf("State (Created:%s, Not Created:%s, Already Exists:%s)", emoji.Sprint(":check_mark_button:"), emoji.Sprint(":cross_mark:"), emoji.Sprint(":white_large_square:")))
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(2)

	var existedFlagKey []string

	for _, flag := range listedFlags {
		existedFlagKey = append(existedFlagKey, strings.ToLower(flag.Name))
	}

	results, err := handler.ExtractFlagsInfo(FSConfig)
	if err != nil {
		return err
	}

	for _, r := range results {
		pathArray := strings.Split(r.File, "/")
		for _, analyzedFlag := range r.Results {

			var flagRequest models.Flag
			var flagResponse models.Flag

			if slices.Contains(existedFlagKey, strings.ToLower(analyzedFlag.FlagKey)) {
				flagAlreadyExistLen += 1
				tbl.AddRow(analyzedFlag.FlagKey, analyzedFlag.FlagType, analyzedFlag.FlagDefaultValue, fmt.Sprintf("%s:%d", pathArray[len(pathArray)-1], analyzedFlag.LineNumber), emoji.Sprint(":white_large_square:"))
				continue
			}

			if analyzedFlag.FlagType == "unknown" {
				flagNotCreatedLen += 1
				flagKeyNotCreated = append(flagKeyNotCreated, analyzedFlag.FlagKey)
				tbl.AddRow(analyzedFlag.FlagKey, analyzedFlag.FlagType, analyzedFlag.FlagDefaultValue, fmt.Sprintf("%s:%d", pathArray[len(pathArray)-1], analyzedFlag.LineNumber), emoji.Sprint(":cross_mark:")+"reason: Unknown type and no default value")
				continue
			}

			if analyzedFlag.FlagType == "boolean" {
				flagRequest = models.Flag{
					Name:        analyzedFlag.FlagKey,
					Type:        analyzedFlag.FlagType,
					Description: "flag created by CLI",
					Source:      "cli",
				}
			} else {
				flagRequest = models.Flag{
					Name:         analyzedFlag.FlagKey,
					Type:         analyzedFlag.FlagType,
					DefaultValue: analyzedFlag.FlagDefaultValue,
					Description:  "flag created by CLI",
					Source:       "cli",
				}
			}

			flagRequestJSON, err_ := json.Marshal(flagRequest)
			if err_ != nil {
				return err_
			}

			createdFlag, errCreatedFlag := httprequest.HTTPCreateFlag(string(flagRequestJSON))

			if errCreatedFlag != nil {
				return errCreatedFlag
			}

			err_json := json.Unmarshal(createdFlag, &flagResponse)

			if err_json != nil {
				return err_json
			}

			if flagResponse.Id != "" {
				flagCreatedLen += 1
				existedFlagKey = append(existedFlagKey, strings.ToLower(analyzedFlag.FlagKey))
				tbl.AddRow(analyzedFlag.FlagKey, analyzedFlag.FlagType, analyzedFlag.FlagDefaultValue, fmt.Sprintf("%s:%d", pathArray[len(pathArray)-1], analyzedFlag.LineNumber), emoji.Sprint(":check_mark_button:"))
			}

		}
	}

	totalFlag := flagCreatedLen + flagAlreadyExistLen + flagNotCreatedLen
	if totalFlag == 0 {
		tbl.AddRow("No flag found")
	}

	tbl.Print()

	summaryTableFlagCreated(flagCreatedLen, flagNotCreatedLen, flagAlreadyExistLen)

	if flagNotCreatedLen != 0 {
		fmt.Fprintf(cmd.OutOrStdout(), "\n%sTips: To create these flags use these commands: \n", emoji.Sprint(":bulb:"))
		for _, flagKey := range flagKeyNotCreated {
			fmt.Fprintf(cmd.OutOrStdout(), "flagship flag create --data-raw '{\"name\": \"%s\",\"type\":\"<TYPE>\",\"description\":\"<DESCRIPTION>\",\"source\":\"cli\"}'\n", flagKey)
		}
	}

	return nil
}

// CreateCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Analyze your codebase and automatically create flags detected",
	Long:  `Analyze your codebase and automatically create flags detected to Flagship platform`,
	PreRun: func(cmd *cobra.Command, args []string) {
		PreRunConfiguration()
	},
	Run: func(cmd *cobra.Command, args []string) {
		listedExistingFlags, errListFlag := httprequest.HTTPListFlag()
		if errListFlag != nil {
			log.Fatalf("error occurred when listing existing flag: %s", errListFlag)
		}

		err := flagCreatedTable(cmd, listedExistingFlags)
		if err != nil {
			log.Fatalf("error occurred in created flag table: %s", err)
		}
	},
}

func init() {
	FlagCmd.AddCommand(createCmd)
}
