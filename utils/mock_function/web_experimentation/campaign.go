package web_experimentation

import (
	"net/http"
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestCampaign = models.CampaignWE{
	Id:                 100000,
	Name:               "testCampaignName",
	Description:        "testCampaignDescription",
	Type:               "ab",
	GlobalCodeCampaign: "console.log(\"Hello World!\")",
}

var TestCampaign1 = models.CampaignWE{
	Id:                 100001,
	Name:               "testCampaignName1",
	Description:        "testCampaignDescription1",
	Type:               "ab",
	GlobalCodeCampaign: "console.log(\"Hello Earth!\")",
}

var TestCampaignList = []models.CampaignWE{
	TestCampaign,
	TestCampaign1,
}

func APICampaign() {

	respList := utils.HTTPListResponseWE[models.CampaignWE]{
		Data: TestCampaignList,
		Pagination: utils.Pagination{
			Total:      1,
			Pages:      2,
			MaxPerPage: 10,
			Page:       1,
		},
	}

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestCampaign)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, respList)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestCampaign)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}
