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
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Directory           string
	RepoURL             string
	RepoBranch          string
	NbLineCodeEdges     int
	FilesToExclude      string
	SearchCustomRegex   string
	CustomRegexJsonFile string
	CustomRegexJson     string
	LaunchDarkly        bool
)
var FSConfig *cbaConfig.Config

func PreRunConfiguration() {
	var filesToExcludeArray []string
	var searchCustomRegex string = SearchCustomRegex

	err := json.Unmarshal([]byte(FilesToExclude), &filesToExcludeArray)
	if err != nil {
		log.Fatalf("error occurred when unmarshal: %s", err)
	}

	if CustomRegexJson != "" {
		searchCustomRegex = CustomRegexJson
	}

	if LaunchDarkly {
		searchCustomRegex = "[{\"file_extension\":\".js?\",\"regexes\": [\"variation(?:Detail)?[(](?:\\\\s*[\\\"\\\\'](.*?)[\\\"\\\\']\\\\s*(?:,\\\\s*([\\\"\\\\'].*\\\\s*[^\\\"]*[\\\"\\\\']|[^)]*))?)\\\\s*[)]\"]}]"
	}

	FSConfig = &cbaConfig.Config{
		FlagshipAPIURL:        utils.GetHost(),
		FlagshipAuthAPIURL:    utils.GetHostAuth(),
		FlagshipAPIToken:      viper.GetString("token"),
		FlagshipAccountID:     viper.GetString("account_id"),
		FlagshipEnvironmentID: viper.GetString("account_environment_id"),
		Directory:             Directory,
		RepositoryURL:         RepoURL,
		RepositoryBranch:      RepoBranch,
		NbLineCodeEdges:       NbLineCodeEdges,
		FilesToExclude:        filesToExcludeArray,
		SearchCustomRegex:     searchCustomRegex,
	}
}

// FlagCmd represents the flag command
var FlagCmd = &cobra.Command{
	Use:   "flag [create|list]",
	Short: "Analyze your codebase and detect the usage of Flagship or custom flags",
	Long:  `Analyze your codebase and detect the usage of Flagship or custom flags, in order to synchronize them with your Flag view in the platform`,
	PreRun: func(cmd *cobra.Command, args []string) {
		PreRunConfiguration()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := handler.AnalyzeCode(FSConfig)
		if err != nil {
			log.Fatalf("error occured when analyzing: %s", err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	FlagCmd.PersistentFlags().StringVarP(&Directory, "directory", "", ".", "directory to analyze in your codebase")
	FlagCmd.PersistentFlags().StringVarP(&RepoURL, "repository-url", "", "https://github.com/org/repo", "root URL of your repository, and is used to track the links of the files where your flags are used")
	FlagCmd.PersistentFlags().StringVarP(&RepoBranch, "repository-branch", "", "main", "branch of the code you want to analyse, and is used to track the links of the files where your flags are used")
	FlagCmd.PersistentFlags().IntVarP(&NbLineCodeEdges, "code-edge", "", 1, "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&FilesToExclude, "files-exclude", "", "[\".git\", \".github\", \".vscode\", \".idea\", \".yarn\", \"node_modules\"]", "list of files to exclude in analysis")
	FlagCmd.PersistentFlags().StringVarP(&SearchCustomRegex, "custom-regex", "", "", "regex for the pattern you want to analyze")
	FlagCmd.PersistentFlags().StringVarP(&CustomRegexJsonFile, "custom-regex-json", "", "", "json file that contains the regex for the pattern you want to analyze")
	FlagCmd.PersistentFlags().BoolVarP(&LaunchDarkly, "launch-darkly", "", false, "analyze flags made with launchdarkly (only latest ones)")
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
