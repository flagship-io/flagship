package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

type TestRequest struct {
	Name string `json:"name"`
}

var Auth = common.RequestConfig{
	Username:             "test_auth",
	ClientID:             "client_id",
	ClientSecret:         "client_secret",
	AccountID:            "account_id",
	AccountEnvironmentID: "account_environment_id",
	Token:                "access_token",
}

func SetMock(c *common.ResourceRequest) {
	viper.GetViper().Set("output_format", "json")
	common.Init(Auth)

	r := c
	r.Init(&Auth)
}

func Request() {

	testRequest := TestRequest{
		Name: "TestName",
	}

	testRequest1 := TestRequest{
		Name: "TestName1",
	}

	testRequestList := []TestRequest{testRequest, testRequest1}

	resp := utils.HTTPListResponseFE[TestRequest]{
		Items:             testRequestList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", "serverURLGet",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, testRequest)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "serverURLList",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)
}
