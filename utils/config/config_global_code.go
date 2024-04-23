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

func VariationGlobalCodeDirectoryJS(workingDir, accountID, campaignID, variationID, code string) string {
	gcWorkingDir := CheckGlobalCodeDirectory(workingDir)
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID

	err := os.MkdirAll(variationCodeDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	jsFilePath := variationCodeDir + "/variationGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing JavaScript file: %s", err)
	}
	return jsFilePath
}

func VariationGlobalCodeDirectoryCSS(workingDir, accountID, campaignID, variationID, code string) string {
	gcWorkingDir := CheckGlobalCodeDirectory(workingDir)
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID

	err := os.MkdirAll(variationCodeDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	jsFilePath := variationCodeDir + "/variationGlobalCode.css"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing CSS file: %s", err)
	}
	return jsFilePath
}

func ElementModificationCodeDirectory(workingDir, accountID, campaignID, variationID, elementID, selector string, code []byte) string {
	gcWorkingDir := CheckGlobalCodeDirectory(workingDir)
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID
	elementCodeDir := variationCodeDir + "/" + elementID

	err := os.MkdirAll(elementCodeDir, os.ModePerm)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	jsFilePath := elementCodeDir + "/element.js"

	err = os.WriteFile(jsFilePath, code, os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing JavaScript file: %s", err)
	}
	return jsFilePath
}

func AddHeaderSelectorComment(selector, code string) []byte {
	selectorComment := "/* Selector: " + selector + " */\n"
	headerComment := []byte(selectorComment)

	fileCode := append(headerComment, []byte(code)...)
	return fileCode
}
