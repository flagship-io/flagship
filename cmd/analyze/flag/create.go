/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
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
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Flag", "File", "Created")
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

				var flagRequest models.Flag
				var flagResponse models.Flag

				if slices.Contains(existedFlagKey, strings.ToLower(result.FlagKey)) {
					/* 					log.WithFields(log.Fields{
						"key": result.FlagKey,
					}).Warn("Existing Flag") */
					state := emoji.Sprint(":warning:")
					tbl.AddRow(result.FlagKey, r.File+":"+strconv.Itoa(result.LineNumber), state)
					continue
				}

				if result.FlagType == "unknown" {
					/* 					log.WithFields(log.Fields{
						"key": result.FlagKey,
					}).Error("Type unknown, Flag not created") */
					state := emoji.Sprint(":cross_mark:")
					tbl.AddRow(result.FlagKey, r.File+":"+strconv.Itoa(result.LineNumber), state)
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
					/* 					log.WithFields(log.Fields{
						"id":           flagResponse.Id,
						"key":          flagResponse.Name,
						"type":         flagResponse.Type,
						"defaultValue": flagResponse.DefaultValue,
						"description":  flagResponse.Description,
						"source":       flagResponse.Source,
					}).Info("Created Flag") */
					state := emoji.Sprint(":check_mark_button:")
					tbl.AddRow(result.FlagKey, r.File+":"+strconv.Itoa(result.LineNumber), state)
				}

			}
		}
		tbl.Print()
	},
}

func init() {
	FlagCmd.AddCommand(createCmd)
}
