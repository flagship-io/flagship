/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/kyokomi/emoji/v2"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

func summaryTableFlagListed(flagExistLen, flagNotExistLen int) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	totalFlag := flagExistLen + flagNotExistLen

	summtbl := table.New("\nSummary")
	summtbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	summtbl.AddRow(fmt.Sprintf("Total flags: %d (%d Flag exist %s, %d Flag don't exist%s)", totalFlag, flagExistLen, emoji.Sprint(":check_mark_button:"), flagNotExistLen, emoji.Sprint(":cross_mark:")))
	summtbl.Print()
}

func flagListedTable(listedFlags []models.Flag) error {

	var flagExistLen int = 0
	var flagNotExistLen int = 0

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Flag", "Type", "defaultValue", "File", fmt.Sprintf("Exists ? (%s/%s)", emoji.Sprint(":check_mark_button:"), emoji.Sprint(":cross_mark:")))
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
			if slices.Contains(existedFlagKey, strings.ToLower(analyzedFlag.FlagKey)) {
				flagExistLen += 1
				tbl.AddRow(analyzedFlag.FlagKey, analyzedFlag.FlagType, analyzedFlag.FlagDefaultValue, fmt.Sprintf("%s/%s:%d", pathArray[len(pathArray)-2], pathArray[len(pathArray)-1], analyzedFlag.LineNumber), emoji.Sprint(":check_mark_button:"))
				continue
			}
			flagNotExistLen += 1
			tbl.AddRow(analyzedFlag.FlagKey, analyzedFlag.FlagType, analyzedFlag.FlagDefaultValue, fmt.Sprintf("%s/%s:%d", pathArray[len(pathArray)-2], pathArray[len(pathArray)-1], analyzedFlag.LineNumber), emoji.Sprint(":cross_mark:"))
		}
	}

	totalFlag := flagExistLen + flagNotExistLen

	if totalFlag == 0 {
		tbl.AddRow("No flag found")
	}

	tbl.Print()

	summaryTableFlagListed(flagExistLen, flagNotExistLen)

	return nil
}

// ListCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Analyze your codebase and list flags detected",
	Long:  `Analyze your codebase and list flags detected and check if it exist in Flagship platform`,
	PreRun: func(cmd *cobra.Command, args []string) {
		PreRunConfiguration()
	},
	Run: func(cmd *cobra.Command, args []string) {

		listExistingFlags, errListFlag := httprequest.HTTPListFlag()
		if errListFlag != nil {
			log.Fatalf("error occurred when listing existing flag: %s", errListFlag)
		}

		err := flagListedTable(listExistingFlags)
		if err != nil {
			log.Fatalf("error occurred in listed flag table: %s", err)
		}

	},
}

func init() {
	FlagCmd.AddCommand(listCmd)
}
