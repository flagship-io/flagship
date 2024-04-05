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

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/cobra"
)

var (
	browser  bool
	password string
	totp     string
)

func checkSingleFlag(bool1, bool2 bool) bool {
	count := 0
	if bool1 {
		count++
	}
	if bool2 {
		count++
	}

	return count == 1
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
	Use:   "login [--browser] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>] | [-u <username> | --username=<username>] [--password <password>]",
	Short: "login",
	Long:  `login`,
	Run: func(cmd *cobra.Command, args []string) {
		if !checkSingleFlag(browser, Username != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (browser, username)")
		}

		if browser {
			codeChan := make(chan string)
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

			go func() {
				http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
					handleCallback(w, r, codeChan)
				})

				if err := http.ListenAndServe("127.0.0.1:8010", nil); err != nil {
					log.Fatalf("Error starting callback server: %s", err)
				}
			}()

			code := <-codeChan

			if code != "" {
				authenticationResponse, err := common.HTTPCreateTokenWEAuthorizationCode(clientID, clientSecret, code)
				if err != nil {
					log.Fatalf("error occurred: %s", err)
					return
				}

				if authenticationResponse.AccessToken == "" {
					log.Fatal("Credentials not valid.")
				}
				// Waiting for fix to implemente route to get username "/users/me"

				fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
				return
			}

			fmt.Fprintln(cmd.OutOrStderr(), "Error occurred.")
		}

		if Username != "" {
			existingCredentials, err := config.GetUsernames(utils.WEB_EXPERIMENTATION)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}
			if slices.Contains(existingCredentials, Username) {
				config.SelectAuth(utils.WEB_EXPERIMENTATION, Username)

				fmt.Fprintln(cmd.OutOrStdout(), "Auth changed successfully to "+Username)
				return
			}

			if password == "" || totp == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (password, totp)")
				return
			}
			authenticationResponse, err := common.HTTPCreateTokenWEPassword(utils.CLIENT_ID, utils.CLIENT_SECRET, Username, password, totp)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
				return
			}

			if authenticationResponse.AccessToken == "" {
				log.Fatal("Credentials not valid.")
			}
			config.CreateAuthFile(utils.WEB_EXPERIMENTATION, Username, "", "", authenticationResponse)
			config.SelectAuth(utils.WEB_EXPERIMENTATION, Username)

			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
		}

	},
}

func init() {

	loginCmd.Flags().StringVarP(&ClientID, "client-id", "i", "", "client ID of an auth")
	loginCmd.Flags().StringVarP(&ClientSecret, "client-secret", "s", "", "client secret of an auth")

	loginCmd.Flags().BoolVarP(&browser, "browser", "", false, "Generate link for browser")
	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "username")
	loginCmd.Flags().StringVarP(&password, "password", "", "", "password")
	loginCmd.Flags().StringVarP(&totp, "totp", "", "", "totp")

	AuthCmd.AddCommand(loginCmd)
}

func handleCallback(w http.ResponseWriter, r *http.Request, codeChan chan<- string) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No token found in URL", http.StatusBadRequest)
		os.Exit(0)
		return
	}

	codeChan <- code

	http.Redirect(w, r, "http://abtasty.com", http.StatusSeeOther)

	go func() {
		time.Sleep(5 * time.Second)
		close(codeChan)
	}()
}
