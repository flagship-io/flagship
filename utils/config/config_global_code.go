package config

import (
	"errors"
	"os"
)

func CheckWorkingDirectory(workingDir string) (string, error) {

	if _, err := os.Stat(workingDir); errors.Is(err, os.ErrNotExist) {
		if err != nil {
			return "", err
		}
	}

	return workingDir, nil
}

func CheckGlobalCodeDirectory(workingDir string) (string, error) {
	wd, err := CheckWorkingDirectory(workingDir)
	if err != nil {
		return "", err
	}

	gcWorkingDir := wd + "/abtasty"

	err = os.MkdirAll(gcWorkingDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return gcWorkingDir, nil
}

func AccountGlobalCodeDirectory(workingDir, accountID, code string) (string, error) {
	gcWorkingDir, err := CheckGlobalCodeDirectory(workingDir)
	if err != nil {
		return "", err
	}
	accountCodeDir := gcWorkingDir + "/" + accountID

	err = os.MkdirAll(accountCodeDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsFilePath := accountCodeDir + "/accountGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		return "", err
	}
	return jsFilePath, nil
}

func CampaignGlobalCodeDirectory(workingDir, accountID, campaignID, code string) (string, error) {
	gcWorkingDir, err := CheckGlobalCodeDirectory(workingDir)
	if err != nil {
		return "", err
	}
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID

	err = os.MkdirAll(campaignCodeDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsFilePath := campaignCodeDir + "/campaignGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		return "", err
	}

	return jsFilePath, nil
}

func VariationGlobalCodeDirectoryJS(workingDir, accountID, campaignID, variationID, code string) (string, error) {
	gcWorkingDir, err := CheckGlobalCodeDirectory(workingDir)
	if err != nil {
		return "", err
	}
	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID

	err = os.MkdirAll(variationCodeDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsFilePath := variationCodeDir + "/variationGlobalCode.js"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		return "", err
	}
	return jsFilePath, nil
}

func VariationGlobalCodeDirectoryCSS(workingDir, accountID, campaignID, variationID, code string) (string, error) {
	gcWorkingDir, err := CheckGlobalCodeDirectory(workingDir)
	if err != nil {
		return "", err
	}

	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID

	err = os.MkdirAll(variationCodeDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsFilePath := variationCodeDir + "/variationGlobalCode.css"
	err = os.WriteFile(jsFilePath, []byte(code), os.ModePerm)
	if err != nil {
		return "", err
	}
	return jsFilePath, nil
}

func ElementModificationCodeDirectory(workingDir, accountID, campaignID, variationID, elementID, selector string, code []byte) (string, error) {
	gcWorkingDir, err := CheckGlobalCodeDirectory(workingDir)
	if err != nil {
		return "", err
	}

	accountCodeDir := gcWorkingDir + "/" + accountID
	campaignCodeDir := accountCodeDir + "/" + campaignID
	variationCodeDir := campaignCodeDir + "/" + variationID
	elementCodeDir := variationCodeDir + "/" + elementID

	err = os.MkdirAll(elementCodeDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	jsFilePath := elementCodeDir + "/element.js"

	err = os.WriteFile(jsFilePath, code, os.ModePerm)
	if err != nil {
		return "", err
	}
	return jsFilePath, nil
}

func AddHeaderSelectorComment(selector, code string) []byte {
	selectorComment := "/* Selector: " + selector + " */\n"
	headerComment := []byte(selectorComment)

	fileCode := append(headerComment, []byte(code)...)
	return fileCode
}
