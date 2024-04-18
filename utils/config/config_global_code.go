package config

import (
	"errors"
	"log"
	"os"
)

func CheckWorkingDirectory(workingDir string) string {

	if _, err := os.Stat(workingDir); errors.Is(err, os.ErrNotExist) {
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
	}

	return workingDir
}

func CheckGlobalCodeDirectory(workingDir string) string {

	wd := CheckWorkingDirectory(workingDir)

	gcWorkingDir := wd + "/abtasty"

	err := os.MkdirAll(gcWorkingDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	return gcWorkingDir
}

func AccountGlobalCodeDirectory(workingDir, accountID, code string) string {
	gcWorkingDir := CheckGlobalCodeDirectory(workingDir)
	accountCodeDir := gcWorkingDir + "/" + accountID

	err := os.MkdirAll(accountCodeDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	jsFilePath := accountCodeDir + "/accountGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing JavaScript file: %s", err)
	}
	return jsFilePath
}

func CampaignGlobalCodeDirectory(workingDir, accountID, campaignID, code string) string {
	gcWorkingDir := CheckGlobalCodeDirectory(workingDir)
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID

	err := os.MkdirAll(campaignCodeDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	jsFilePath := campaignCodeDir + "/campaignGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing JavaScript file: %s", err)
	}
	return jsFilePath
}
