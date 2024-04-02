package feature_experimentation

import (
	"net/http"
	"net/url"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestUserList = []models.User{
	{
		Email: "example@abtasty.com",
		Role:  "ADMIN",
	},
	{
		Email: "example1@abtasty.com",
		Role:  "VIEWER",
	},
}

func APIUser() {

	email := "example@abtasty.com"

	resp := utils.HTTPListResponse[models.User]{
		Items:             TestUserList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PUT", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(204, "")
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/users?emails[]="+url.QueryEscape(email),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(204, "")
			return resp, nil
		},
	)
}
