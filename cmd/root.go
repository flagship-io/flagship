/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package cmd

import (
	"os"

	"github.com/flagship-io/flagship/cmd/configuration"
	"github.com/flagship-io/flagship/cmd/feature_experimentation"
	"github.com/flagship-io/flagship/cmd/info"
	"github.com/flagship-io/flagship/cmd/token"

	"github.com/flagship-io/flagship/cmd/version"
	"github.com/flagship-io/flagship/utils/config"
	httprequest_fe "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	cmdToken     string
	outputFormat string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagship",
	Short: "flagship manage your campaigns, project, users etc...",
	Long: `flagship is the main command, used to manage campaigns, projects, users, variation groups and variations
	
	Flagship is a feature flagging platform for modern developers. 
	Separate code deployments from feature releases to accelerate development cycles and mitigate risks.
	
	Complete documentation is available at https://docs.developers.flagship.io/docs/flagship-command-line-interface`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addFeatureExpSubCommandPalettes() {
	rootCmd.AddCommand(configuration.ConfigurationCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(token.TokenCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(feature_experimentation.FeatureExperimentationCmd)
}

func addWebExpSubCommandPalettes() {

}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cmdToken, "token", "t", "", "access token to manage flagship resources")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", config.OutputFormat, "output format for the get and list subcommands for flagship resources. Only 3 format are possible: table, json, json-pretty")
	rootCmd.PersistentFlags().StringVarP(&httprequest_fe.UserAgent, "user-agent", "u", config.DefaultUserAgent, "custom user agent")

	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output-format"))

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file that contains your credentials (default is $HOME/.flagship/credentials.yaml)")

	addFeatureExpSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.MergeInConfig()
		return
	}
	// Find home directory.
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	// Search config in home directory with name ".flagship" (without extension).
	viper.SetConfigFile(homeDir + "/.flagship/configurations/.cli.yaml")
	viper.MergeInConfig()
	if viper.GetString("current_used_configuration") != "" {
		config.ReadConfiguration(viper.GetString("current_used_configuration"))
	}

}
