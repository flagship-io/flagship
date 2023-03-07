/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"log"
	"os"

	cbaConfig "github.com/flagship-io/codebase-analyzer/pkg/config"
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
var FSConfig *cbaConfig.Config

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

		FSConfig = &cbaConfig.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipClientID:      viper.GetString("client_id"),
			FlagshipClientSecret:  viper.GetString("client_secret"),
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

	FlagCmd.PersistentFlags().StringVarP(&Directory, "directory", "", ".", "directory to analyze in your codebase")
	FlagCmd.PersistentFlags().StringVarP(&RepoURL, "repository-url", "", "https://github.com/org/repo", "root URL of your repository, and is used to track the links of the files where your flags are used")
	FlagCmd.PersistentFlags().StringVarP(&RepoBranch, "repository-branch", "", "main", "branch of the code you want to analyse, and is used to track the links of the files where your flags are used")
	FlagCmd.PersistentFlags().IntVarP(&NbLineCodeEdges, "code-edge", "", 1, "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&FilesToExcludes, "files-exclude", "", "[\".git\", \".github\", \".vscode\", \".idea\"]", "list of files to exclude in analysis")
	FlagCmd.PersistentFlags().StringVarP(&SearchCustomRegex, "custom-regex", "", "", "regex for the pattern you want to analyze")
	FlagCmd.PersistentFlags().StringVarP(&CustomRegexJsonFile, "custom-regex-json", "", "", "json file that the regex for the pattern you want to analyze")
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
