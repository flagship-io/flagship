package httprequest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type TestRequest struct {
	Name string `json:"name"`
}

func TestHTTPRequestGet(t *testing.T) {

	var result TestRequest

	testRequest := TestRequest{
		Name: "TestName",
	}
	testRequestJson, _ := json.Marshal(testRequest)

	respBody, err := HTTPRequest(http.MethodGet, "serverURLGet", testRequestJson)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	json.Unmarshal(respBody, &result)

	assert.Equal(t, "TestName", result.Name)
}

func TestHTTPGetItem(t *testing.T) {

	result, err := HTTPGetItem[TestRequest]("serverURLGet")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "TestName", result.Name)
}

func TestHTTPGetAllPages(t *testing.T) {

	result, err := HTTPGetAllPages[TestRequest]("serverURLList")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, "TestName", result[0].Name)
	assert.Equal(t, "TestName1", result[1].Name)
}

func TestRegenerateToken(t *testing.T) {

	regenerateToken("test_configuration")

	assert.Equal(t, viper.IsSet("token"), true)
	assert.Equal(t, viper.GetString("token"), "access_token")
}
