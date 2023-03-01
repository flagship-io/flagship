/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"log"
	"os"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Directory           string
	RepoURL             string
	RepoBranch          string
	NbLineCodeEdges     int
	FilesToExcludes     string
	SearchCustomRegex   string
	CustomRegexJsonFile string
	CustomRegexJson     string
)
var FSConfig *config.Config

// FlagCmd represents the flag command
var FlagCmd = &cobra.Command{
	Use:   "flag [create|list]",
	Short: "Analyze your codebase and detect the usage of Flagship or custom flags",
	Long:  `Analyze your codebase and detect the usage of Flagship or custom flags, in order to synchronize them with your Flag view in the platform`,
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
		err := handler.AnalyzeCode(FSConfig)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	FlagCmd.PersistentFlags().StringVarP(&Directory, "directory", "", ".", "directory to analyze")
	FlagCmd.PersistentFlags().StringVarP(&RepoURL, "repository-url", "", "https://gitlab.com/org/repo", "repository URL")
	FlagCmd.PersistentFlags().StringVarP(&RepoBranch, "repository-branch", "", "main", "repository branch")
	FlagCmd.PersistentFlags().IntVarP(&NbLineCodeEdges, "code-edge", "", 1, "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&FilesToExcludes, "file-excludes", "", "[\".git\", \".github\", \".vscode\", \".idea\"]", "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&SearchCustomRegex, "custom-regex", "", "", "custom regex")
	FlagCmd.PersistentFlags().StringVarP(&CustomRegexJsonFile, "json", "", "", "custom regex in json")
}

func initConfig() {

	if CustomRegexJsonFile != "" {
		bytes, err := os.ReadFile(CustomRegexJsonFile)

		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}

		CustomRegexJson = string(bytes)

	}
}
