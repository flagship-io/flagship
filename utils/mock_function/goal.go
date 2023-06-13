package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
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
	config.SetViperMock()

	resp := utils.HTTPListResponse[models.Goal]{
		Items:             TestGoalList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoal)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoal)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestGoalEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+TestGoal.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)
}
