/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"time"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/cobra"
)

var (
	credentialsFile string
	browser         bool
)

var (
	code string
)

func checkSingleFlag(bool1, bool2, bool3 bool) bool {
	return (bool1 && !bool2 && !bool3) || (!bool1 && bool2 && !bool3) || (!bool1 && !bool2 && bool3)
}
func openLink(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported operating system")
	}
	err := cmd.Run()
	return err
}

// createCmd represents the create command
var loginCmd = &cobra.Command{
	Use:   "login [--browser] | [--credential-file] | [--inputs] [-u <username> | --username=<username>] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>]",
	Short: "login",
	Long:  `login`,
	Run: func(cmd *cobra.Command, args []string) {
		if !checkSingleFlag(browser, credentialsFile != "", Username != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (browser, inputs, credential-file)")
		}

		if browser {
			clientID := utils.CLIENT_ID
			clientSecret := utils.CLIENT_SECRET

			if ClientID != "" {
				clientID = ClientID
			}

			if ClientSecret != "" {
				clientSecret = ClientSecret
			}

			var url = fmt.Sprintf("https://auth.abtasty.com/authorize?client_id=%s&client_secret=%s&redirect_uri=http://localhost:8010/auth/callback", clientID, clientSecret)

			if err := openLink(url); err != nil {
				log.Fatalf("Error opening link: %s", err)
			}
			http.HandleFunc("/auth/callback", handleCallback)
			if err := http.ListenAndServe("127.0.0.1:8010", nil); err != nil {
				log.Fatalf("Error starting callback server: %s", err)
			}
			fmt.Println(code)

			authenticationResponse, err := common.HTTPCreateTokenWE(clientID, clientSecret, code)
			if err != nil {
				log.Fatalf("%s", err)
				return
			}

			if authenticationResponse.AccessToken == "" {
				log.Fatal("client_id or client_secret not valid")
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Token generated successfully")

		}

		if credentialsFile != "" {
			authenticationResponse := models.TokenResponse{}
			config.ReadCredentialsFromFile(credentialsFile)
			config.CreateCredentialsFile(utils.FEATURE_EXPERIMENTATION, Username, ClientID, ClientSecret, authenticationResponse)
			config.SelectCredentials(utils.FEATURE_EXPERIMENTATION, Username)
			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
			return
		}

		if Username != "" {
			existingCredentials, err := config.GetUsernames(utils.FEATURE_EXPERIMENTATION)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}
			if slices.Contains(existingCredentials, Username) {
				config.SelectCredentials(utils.FEATURE_EXPERIMENTATION, Username)
				config.SetAccountID(utils.FEATURE_EXPERIMENTATION, AccountId)

				fmt.Fprintln(cmd.OutOrStdout(), "Auth changed successfully to "+Username)
				return
			}

			if ClientID == "" && ClientSecret == "" && AccountId == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (username, client ID, client secret, account id)")
				return
			}
			authenticationResponse, err := common.HTTPCreateTokenFE(ClientID, ClientSecret, AccountId)
			if err != nil {
				log.Fatalf("%s", err)
				return
			}

			if authenticationResponse.AccessToken == "" {
				log.Fatal("client_id or client_secret not valid")
			}
			config.CreateCredentialsFile(utils.FEATURE_EXPERIMENTATION, Username, ClientID, ClientSecret, authenticationResponse)
			config.SelectCredentials(utils.FEATURE_EXPERIMENTATION, Username)
			config.SetAccountID(utils.FEATURE_EXPERIMENTATION, AccountId)

			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
		}

	},
}

func init() {

	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "configuration name")
	loginCmd.Flags().StringVarP(&ClientID, "client-id", "i", "", "client ID of a configuration")
	loginCmd.Flags().StringVarP(&ClientSecret, "client-secret", "s", "", "client secret of a configuration")
	loginCmd.Flags().StringVarP(&AccountId, "account-id", "a", "", "account id of a configuration")

	loginCmd.Flags().StringVarP(&credentialsFile, "credential-file", "p", "", "config file to create")
	loginCmd.Flags().BoolVarP(&browser, "browser", "", false, "Generate link for browser")

	AuthCmd.AddCommand(loginCmd)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("code"))
	code = r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No token found in URL", http.StatusBadRequest)
		os.Exit(0)
		return
	}

	http.Redirect(w, r, "http://abtasty.com", http.StatusSeeOther)

	fmt.Println("code received:", code)

	go func() {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()
}
