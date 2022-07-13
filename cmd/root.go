/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/Chadiii/flagship-mock/cmd/authorization"
	"github.com/Chadiii/flagship-mock/cmd/campaign"
	"github.com/Chadiii/flagship-mock/cmd/panic"
	"github.com/Chadiii/flagship-mock/cmd/project"
	"github.com/Chadiii/flagship-mock/cmd/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile                string
	Token                  string
	Account_id             string
	Account_environment_id string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagship-mock",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

func addSubCommandPalettes() {
	rootCmd.AddCommand(campaign.CampaignCmd)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(authorization.AuthorizationCmd)
	//rootCmd.AddCommand(variation_group.VariationGroupCmd)
	rootCmd.AddCommand(panic.PanicCmd)
	rootCmd.AddCommand(user.UserCmd)
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Token, "token", "", "", "authorization token")
	rootCmd.PersistentFlags().StringVarP(&Account_id, "account_id", "", "", "account id")
	rootCmd.PersistentFlags().StringVarP(&Account_environment_id, "account_environment_id", "", "", "account env id")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("account_id", rootCmd.PersistentFlags().Lookup("account_id"))
	viper.BindPFlag("account_environment_id", rootCmd.PersistentFlags().Lookup("account_environment_id"))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship-mock" (without extension).
		viper.SetConfigFile(workingDir + "/config.yaml")
	}

	// read in environment variables that match

	// If a config file is found, read it in.
	viper.MergeInConfig()
}
