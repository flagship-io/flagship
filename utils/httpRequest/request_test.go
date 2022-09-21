package httprequest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type TestRequest struct {
	Name string `json:"name"`
}

type HTTPListResponse[T any] struct {
	Items             []T `json:"items"`
	CurrentItemsCount int `json:"current_items_count"`
	CurrentPage       int `json:"current_page"`
	TotalCount        int `json:"total_count"`
	ItemsPerPage      int `json:"items_per_page"`
	LastPage          int `json:"last_page"`
}

func New(exit Func) *Exit {
	return &Exit{exit: exit}
}

type Func func(int)

type Exit struct {
	exit   Func
	status int
}

func (e *Exit) Exit(code int) {
	if e != nil {
		e.status = code
		e.exit(code)
	} else {
		os.Exit(code)
	}
}

func ViperNotSet(t *testing.T) {
	exiter := New(func(int) {})
	exiter.Exit(1)

	if !viper.IsSet("account_id") {
		assert.Equal(t, exiter.status, 1)
	}

	if !viper.IsSet("account_environment_id") {
		assert.Equal(t, exiter.status, 1)
	}

	if !viper.IsSet("client_id") {
		assert.Equal(t, exiter.status, 1)
	}

	if !viper.IsSet("client_secret") {
		assert.Equal(t, exiter.status, 1)
	}

	if !viper.IsSet("token") {
		assert.Equal(t, exiter.status, 1)
	}

	viper.Set("account_id", "account_id")
	viper.Set("account_environment_id", "account_environment_id")
	viper.Set("client_id", "client_id")
	viper.Set("client_secret", "client_secret")
	viper.Set("token", "token")
}

func TestHTTPRequestGet(t *testing.T) {

	ViperNotSet(t)

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

	ViperNotSet(t)

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

	ViperNotSet(t)

	testRequest1 := TestRequest{
		Name: "TestName1",
	}
	testRequest2 := TestRequest{
		Name: "TestName2",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		resp := HTTPListResponse[TestRequest]{
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
