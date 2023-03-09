package mockfunction

import (
	"net/http"
	"net/url"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
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
	config.SetViper()

	email := "example@abtasty.com"

	resp := utils.HTTPListResponse[models.User]{
		Items:             TestUserList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PUT", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(204, "")
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+url.QueryEscape(email),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(204, "")
			return resp, nil
		},
	)
}
