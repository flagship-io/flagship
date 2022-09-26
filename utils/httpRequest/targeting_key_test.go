package httprequest

import (
	"net/http"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHTTPGetTargetingKey(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testTargetingKey := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName",
		Type:        "string",
		Description: "testTargetingKeyDescription",
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKey)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetTargetingKey("testTargetingKeyID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody.ID)
	assert.Equal(t, "testTargetingKeyName", respBody.Name)
}

func TestHTTPListTargetingKey(t *testing.T) {

	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testTargetingKeyList := []models.TargetingKey{
		{
			ID:          "testTargetingKeyID",
			Name:        "testTargetingKeyName",
			Type:        "string",
			Description: "testTargetingKeyDescription",
		},
		{
			ID:          "testTargetingKeyID1",
			Name:        "testTargetingKeyName1",
			Type:        "string",
			Description: "testTargetingKeyDescription1",
		},
	}

	resp := HTTPListResponse[models.TargetingKey]{
		Items:             testTargetingKeyList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListTargetingKey()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody[0].ID)
	assert.Equal(t, "testTargetingKeyName", respBody[0].Name)

	assert.Equal(t, "testTargetingKeyID1", respBody[1].ID)
	assert.Equal(t, "testTargetingKeyName1", respBody[1].Name)
}

func TestHTTPCreateTargetingKey(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testTargetingKey := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName",
		Type:        "string",
		Description: "testTargetingKeyDescription",
	}

	data := "{\"name\":\"testTargetingKeyName\", \"type\":\"string\", \"description\":\"testTargetingKeyDescription\"}"

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKey)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateTargetingKey(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription\"}"), respBody)
}

func TestHTTPEditTargetingKey(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testTargetingKey := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName1",
		Type:        "string",
		Description: "testTargetingKeyDescription1",
	}

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKey)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditTargetingKey(testTargetingKey.ID, testTargetingKey.Name)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName1\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription1\"}"), respBody)
}

func TestHTTPDeleteTargetingKey(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testTargetingKey := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName",
		Type:        "string",
		Description: "testTargetingKeyDescription",
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteTargetingKey("testTargetingKeyID")

	assert.Nil(t, err)
}
