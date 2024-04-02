package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestFlag = models.Flag{
	Id:          "testFlagID",
	Name:        "testFlagName",
	Type:        "string",
	Description: "testFlagDescription",
	Source:      "cli",
}

var TestFlag1 = models.Flag{
	Id:          "testFlagID1",
	Name:        "testFlagName1",
	Type:        "string",
	Description: "testFlagDescription1",
	Source:      "cli",
}

var TestFlagEdit = models.Flag{
	Id:          "testFlagID",
	Name:        "testFlagName1",
	Type:        "string",
	Description: "testFlagDescription1",
	Source:      "cli",
}

var TestFlagList = []models.Flag{
	TestFlag,
	TestFlag1,
}

func APIFlag() {

	resp := utils.HTTPListResponse[models.Flag]{
		Items:             TestFlagList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/flags/"+TestFlag.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestFlag)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestFlag)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/flags/"+TestFlag.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestFlagEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/flags/"+TestFlag.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}
