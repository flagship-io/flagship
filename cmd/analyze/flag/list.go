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
	Short: "List your flags",
	Long:  `List your flags present in the directory`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var FilesToExcludes_ []string

		err := json.Unmarshal([]byte(FilesToExcludes), &FilesToExcludes_)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
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
			FilesToExcludes:       FilesToExcludes_,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var flagLen int = 0
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Flag", "File", "Exists")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

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
				flagLen += 1
				if slices.Contains(existedFlagKey, strings.ToLower(result.FlagKey)) {
					state := emoji.Sprint(":check_mark_button:")
					tbl.AddRow(result.FlagKey, r.File+":"+strconv.Itoa(result.LineNumber), state)
					continue
				}
				state := emoji.Sprint(":cross_mark:")
				tbl.AddRow(result.FlagKey, r.File+":"+strconv.Itoa(result.LineNumber), state)
			}
		}

		if flagLen == 0 {
			tbl.AddRow("No flag found")
		}
		tbl.AddRow("Total flags: " + strconv.Itoa(flagLen) + " found")
		tbl.Print()
	},
}

func init() {
	FlagCmd.AddCommand(listCmd)
}
