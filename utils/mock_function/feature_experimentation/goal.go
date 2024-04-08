package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestGoal = models.Goal{
	Id:       "testGoalID",
	Label:    "testGoalLabel",
	Type:     "screenview",
	Operator: "contains",
	Value:    "VIP",
}

var TestGoal1 = models.Goal{
	Id:       "testGoalID1",
	Label:    "testGoalLabel1",
	Type:     "pageview",
	Operator: "exact",
	Value:    "notVIP",
}

var TestGoalEdit = models.Goal{
	Id:       "testGoalID",
	Label:    "testGoalLabel1",
	Type:     "screenview",
	Operator: "contains",
	Value:    "VIP1",
}

var TestGoalList = []models.Goal{
	TestGoal,
	TestGoal1,
}

func APIGoal() {

	resp := utils.HTTPListResponseFE[models.Goal]{
		Items:             TestGoalList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoal)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/goals",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/goals",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoal)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoalEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}
