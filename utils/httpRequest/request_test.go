package httprequest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type TestRequest struct {
	Name string `json:"name"`
}

func TestHTTPRequestGet(t *testing.T) {

	config.ViperNotSet(t)

	var result TestRequest

	testRequest := TestRequest{
		Name: "TestName",
	}
	testRequestJson, _ := json.Marshal(testRequest)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			t.Errorf("want %s, got %s", "GET", req.Method)
		}
		_, err := rw.Write(testRequestJson)
		assert.Nil(t, err)
	}))

	defer server.Close()

	respBody, err := HTTPRequest(http.MethodGet, server.URL, testRequestJson)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	json.Unmarshal(respBody, &result)

	assert.Equal(t, "TestName", result.Name)
}

func TestHTTPGetItem(t *testing.T) {

	config.ViperNotSet(t)

	testRequest := TestRequest{
		Name: "TestName",
	}
	testRequestJson, _ := json.Marshal(testRequest)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			t.Errorf("want %s, got %s", "GET", req.Method)
		}
		_, err := rw.Write(testRequestJson)
		assert.Nil(t, err)
	}))
	defer server.Close()

	result, err := HTTPGetItem[TestRequest](server.URL)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "TestName", result.Name)
}

func TestHTTPGetAllPages(t *testing.T) {

	config.ViperNotSet(t)

	testRequest1 := TestRequest{
		Name: "TestName1",
	}
	testRequest2 := TestRequest{
		Name: "TestName2",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		resp := utils.HTTPListResponse[TestRequest]{
			Items:             []TestRequest{testRequest1, testRequest2},
			CurrentItemsCount: 2,
			CurrentPage:       1,
			TotalCount:        2,
			ItemsPerPage:      10,
			LastPage:          1,
		}

		testRequestsJson, err := json.Marshal(resp)

		assert.Nil(t, err)

		_, err1 := rw.Write(testRequestsJson)
		assert.Nil(t, err1)
	}))

	defer server.Close()

	result, err := HTTPGetAllPages[TestRequest](server.URL)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "TestName1", result[0].Name)
	assert.Equal(t, "TestName2", result[1].Name)
}

func TestRegenerateToken(t *testing.T) {

	regenerateToken("credentialsTest.yaml")

	assert.Equal(t, viper.IsSet("token"), true)
	assert.Equal(t, viper.GetString("token"), "access_token")
}
