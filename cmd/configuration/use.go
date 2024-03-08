/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package configuration

import (
	"fmt"
	"log"
	"slices"

	"github.com/flagship-io/flagship/utils/config"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	grantType  string
	scope      string
	expiration int
)

// useCmd represents use command
var useCmd = &cobra.Command{
	Use:   "use [-n <name> | --name=<name>] [--grant-type=<grant-type>] [--scope=<scope>] [--expiration=<expiration>]",
	Short: "Use a configuration",
	Long:  `Use an already created configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		existingConfigurationsName, err := config.GetConfigurationsName()
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
		if !slices.Contains(existingConfigurationsName, ConfigurationName) {
			fmt.Fprintln(cmd.OutOrStdout(), "Configuration doesn't exists")
			return
		}

		config.SelectConfiguration(ConfigurationName)
		config.SetOptionalsDefault(grantType, scope, expiration)

		token, err := httprequest.HTTPCreateToken(viper.GetString("client_id"), viper.GetString("client_secret"), grantType, scope, expiration)
		if err != nil {
			log.Fatalf("%s", err)
			return
		}

		if token == "" {
			log.Fatal("client_id or client_secret not valid")
			return
		} else {
			config.WriteToken(ConfigurationName, token)
			fmt.Fprintln(cmd.OutOrStdout(), "Token generated successfully")
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Configuration selected successfully")
	},
}

func init() {
	useCmd.Flags().StringVarP(&ConfigurationName, "name", "n", "", "name of the configuration you want to display")
	useCmd.Flags().StringVarP(&grantType, "grant-type", "", config.GrantType, "grant type of the token, DEFAULT value is client_credentials")
	useCmd.Flags().StringVarP(&scope, "scope", "", config.Scope, "scope of the token, DEFAULT value is *")
	useCmd.Flags().IntVarP(&expiration, "expiration", "", config.Expiration, "expiration time in second of the token, DEFAULT value is 86400")

	if err := useCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	ConfigurationCmd.AddCommand(useCmd)
}
