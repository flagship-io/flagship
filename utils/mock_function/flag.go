package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

func APIFlag() {
	config.SetViper()

	testFlag := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName",
		Type:        "string",
		Description: "testFlagDescription",
		Source:      "manual",
	}

	testFlag1 := models.Flag{
		ID:          "testFlagID1",
		Name:        "testFlagName1",
		Type:        "string",
		Description: "testFlagDescription1",
		Source:      "manual",
	}

	testFlagEdit := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName1",
		Type:        "string",
		Description: "testFlagDescription1",
		Source:      "manual",
	}

	testFlagList := []models.Flag{
		testFlag,
		testFlag1,
	}

	resp := utils.HTTPListResponse[models.Flag]{
		Items:             testFlagList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	testFlagUsageList := []models.FlagUsage{
		{
			ID:                "testFlagUsageID",
			FlagKey:           "isVIP",
			Repository:        "flagship-cli",
			FilePath:          "https://github.com/flagship-io/flagship-cli",
			Branch:            "main",
			Line:              "Line116",
			CodeLineHighlight: "codeLineHighlight",
			Code:              "code",
		},
	}

	respUsage := utils.HTTPListResponse[models.FlagUsage]{
		Items:             testFlagUsageList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlag)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/flags_usage",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, respUsage)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlag)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlagEdit)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}
