package httprequest

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHTTPListUsers(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testUserList := []models.User{
		{
			Email: "example@abtasty.com",
			Role:  "ADMIN",
		},
		{
			Email: "example1@abtasty.com",
			Role:  "VIEWER",
		},
	}

	resp := utils.HTTPListResponse[models.User]{
		Items:             testUserList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListUsers()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "example@abtasty.com", respBody[0].Email)
	assert.Equal(t, "ADMIN", respBody[0].Role)

	assert.Equal(t, "example1@abtasty.com", respBody[1].Email)
	assert.Equal(t, "VIEWER", respBody[1].Role)
}

func TestHTTPBatchUpdateUsers(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := "[{\"email\":\"example@abtasty.com\",\"role\":\"ADMIN\"},{\"email\":\"example1@abtasty.com\",\"role\":\"VIEWER\"}]"

	httpmock.RegisterResponder("PUT", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(204, "")
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	_, err := HTTPBatchUpdateUsers(data)

	assert.Nil(t, err)
}

func TestHTTPDeleteUser(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	email := "example@abtasty.com"

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+url.QueryEscape(email),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(204, "")
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	err := HTTPDeleteUsers(email)

	assert.Nil(t, err)
}
