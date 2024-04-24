package feature_experimentation

import (
	"net/http"

	models_fe "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestAccountEnvironment = models_fe.AccountEnvironmentFE{
	Id:               "account_environment_id",
	Environment:      "account_environment_name",
	IsMain:           true,
	Panic:            false,
	SingleAssignment: false,
}

var TestAccountEnvironment1 = models_fe.AccountEnvironmentFE{
	Id:               "account_environment_id1",
	Environment:      "account_environment_name1",
	IsMain:           true,
	Panic:            false,
	SingleAssignment: false,
}

var TestAccountEnvironmentList = []models_fe.AccountEnvironmentFE{
	TestAccountEnvironment, TestAccountEnvironment1,
}

func APIAccountEnvironment() {

	resp := utils.HTTPListResponseFE[models_fe.AccountEnvironmentFE]{
		Items:             TestAccountEnvironmentList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+TestAccountEnvironment.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestAccountEnvironment)
			return resp, nil
		},
	)
}
