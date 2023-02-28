/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/kyokomi/emoji/v2"
	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

// CreateCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create analyzed flags",
	Long:  `create analyzed flag present in the directory`,
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
			SearchCustomRegex:     SearchCustomRegex,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var flagCreatedLen int = 0
		var flagNotCreatedLen int = 0
		var flagAlreadyExistLen int = 0
		var flagKeyNotCreated []string

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Flag", "Type", "defaultValue", "File", "State (Created:"+emoji.Sprint(":check_mark_button:")+", Not Created:"+emoji.Sprint(":cross_mark:")+", Already Exists:"+emoji.Sprint(":white_large_square:")+")")
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

				var flagRequest models.Flag
				var flagResponse models.Flag

				if slices.Contains(existedFlagKey, strings.ToLower(result.FlagKey)) {
					flagAlreadyExistLen += 1
					tbl.AddRow(result.FlagKey, result.FlagType, result.FlagDefaultValue, r.File+":"+strconv.Itoa(result.LineNumber), emoji.Sprint(":white_large_square:"))
					continue
				}

				if result.FlagType == "unknown" {
					flagNotCreatedLen += 1
					flagKeyNotCreated = append(flagKeyNotCreated, result.FlagKey)
					tbl.AddRow(result.FlagKey, result.FlagType, result.FlagDefaultValue, r.File+":"+strconv.Itoa(result.LineNumber), emoji.Sprint(":cross_mark:")+"reason: Unknown type and no default value")
					continue
				}

				if result.FlagType == "boolean" {
					flagRequest = models.Flag{
						Name:        result.FlagKey,
						Type:        result.FlagType,
						Description: "flag created by CLI",
						Source:      "codebase_analyzer",
					}
				} else {
					flagRequest = models.Flag{
						Name:         result.FlagKey,
						Type:         result.FlagType,
						DefaultValue: result.FlagDefaultValue,
						Description:  "flag created by CLI",
						Source:       "codebase_analyzer",
					}
				}

				flagRequestJSON, err_ := json.Marshal(flagRequest)
				if err_ != nil {
					log.Fatalf("error occurred: %s", err)
				}

				createdFlag, errCreatedFlag := httprequest.HTTPCreateFlag(string(flagRequestJSON))

				if errCreatedFlag != nil {
					log.Fatalf("error occurred: %v", err)
				}

				err_json := json.Unmarshal(createdFlag, &flagResponse)

				if err_json != nil {
					log.Fatalf("error occurred: %v", err)
				}

				if flagResponse.Id != "" {
					flagCreatedLen += 1
					tbl.AddRow(result.FlagKey, result.FlagType, result.FlagDefaultValue, r.File+":"+strconv.Itoa(result.LineNumber), emoji.Sprint(":check_mark_button:"))
				}

			}
		}

		totalFlag := flagCreatedLen + flagAlreadyExistLen + flagNotCreatedLen
		if totalFlag == 0 {
			tbl.AddRow("No flag found")
		}

		summtbl.AddRow("Total flags: " + strconv.Itoa(totalFlag) + " (" + strconv.Itoa(flagCreatedLen) + " Flag created " + emoji.Sprint(":check_mark_button:") + ", " + strconv.Itoa(flagNotCreatedLen) + " Flag not created" + emoji.Sprint(":cross_mark:") + ", " + strconv.Itoa(flagAlreadyExistLen) + " Flag that already exist" + emoji.Sprint(":white_large_square:") + ")")

		tbl.Print()

		summtbl.Print()

		if flagNotCreatedLen != 0 {
			fmt.Fprintf(cmd.OutOrStdout(), "\n%sTips: To create these flags use these commands: \n", emoji.Sprint(":bulb:"))
			for _, flagKey := range flagKeyNotCreated {
				fmt.Fprintf(cmd.OutOrStdout(), "flagship flag create --data-raw '{\"name\": \"%s\",\"type\":\"<TYPE>\",\"description\":\"<DESCRIPTION>\",\"source\":\"cli\"}'\n", flagKey)
			}
		}
	},
}

func init() {
	FlagCmd.AddCommand(createCmd)
}
