/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package analyse

import (
	"log"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	directory   string
	repoURL     string
	repoBranch  string
	excluseFile string
	customRegex string
)

// analyseCmd represents the analyse command
var AnalyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Manage your flags",
	Long:  `Manage your flags in your account`,
	Run: func(cmd *cobra.Command, args []string) {

		err := handler.AnalyzeCode(&config.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipAccountID:     viper.GetString("account_id"),
			FlagshipEnvironmentID: viper.GetString("account_environment_id"),
			Directory:             ".",
			RepositoryURL:         "https://github.com/Chadiii/react-app-with-flagship-sdk",
			RepositoryBranch:      "main",
			NbLineCodeEdges:       1,
		})

		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}
