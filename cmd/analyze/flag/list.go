/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/kyokomi/emoji/v2"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// ListCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Analyze your codebase and list flags detected",
	Long:  `Analyze your codebase and list flags detected and check if it exist in Flagship platform`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var filesToExcludes_ []string
		var searchCustomRegex string = SearchCustomRegex

		err := json.Unmarshal([]byte(FilesToExcludes), &filesToExcludes_)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		if CustomRegexJson != "" {
			searchCustomRegex = CustomRegexJson
		}

		FSConfig = &config.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipAccountID:     viper.GetString("account_id"),
			FlagshipEnvironmentID: viper.GetString("account_environment_id"),
			Directory:             Directory,
			RepositoryURL:         RepoURL,
			RepositoryBranch:      RepoBranch,
			NbLineCodeEdges:       NbLineCodeEdges,
			FilesToExcludes:       filesToExcludes_,
			SearchCustomRegex:     searchCustomRegex,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var flagExistLen int = 0
		var flagNotExistLen int = 0

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Flag", "Type", "defaultValue", "File", "Exists ? ("+emoji.Sprint(":check_mark_button:")+"/"+emoji.Sprint(":cross_mark:")+")")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		summtbl := table.New("\nSummary")
		summtbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		var existedFlagKey []string

		listedFlags, errListFlag := httprequest.HTTPListFlag()
		if errListFlag != nil {
			log.Fatalf("error occurred: %v", errListFlag)
		}

		for _, flag := range listedFlags {
			existedFlagKey = append(existedFlagKey, strings.ToLower(flag.Name))
		}

		results, err := handler.ExtractFlagsInfo(FSConfig)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}

		for _, r := range results {
			for _, result := range r.Results {
				if slices.Contains(existedFlagKey, strings.ToLower(result.FlagKey)) {
					flagExistLen += 1
					tbl.AddRow(result.FlagKey, result.FlagType, result.FlagDefaultValue, r.File+":"+strconv.Itoa(result.LineNumber), emoji.Sprint(":check_mark_button:"))
					continue
				}
				flagNotExistLen += 1
				tbl.AddRow(result.FlagKey, result.FlagType, result.FlagDefaultValue, r.File+":"+strconv.Itoa(result.LineNumber), emoji.Sprint(":cross_mark:"))
			}
		}

		totalFlag := flagExistLen + flagNotExistLen
		if totalFlag == 0 {
			tbl.AddRow("No flag found")
		}

		tbl.Print()

		summtbl.AddRow("Total flags: " + strconv.Itoa(totalFlag) + " (" + strconv.Itoa(flagExistLen) + " Flag exist " + emoji.Sprint(":check_mark_button:") + ", " + strconv.Itoa(flagNotExistLen) + " Flag don't exist" + emoji.Sprint(":cross_mark:") + ")")
		summtbl.Print()
	},
}

func init() {
	FlagCmd.AddCommand(listCmd)
}
