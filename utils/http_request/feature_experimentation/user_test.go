package feature_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var userRequester = UserRequester{&common.ResourceRequest{AccountID: "account_id", AccountEnvironmentID: "account_environment_id"}}

func TestHTTPListUsers(t *testing.T) {

	respBody, err := userRequester.HTTPListUsers()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "example@abtasty.com", respBody[0].Email)
	assert.Equal(t, "ADMIN", respBody[0].Role)

	assert.Equal(t, "example1@abtasty.com", respBody[1].Email)
	assert.Equal(t, "VIEWER", respBody[1].Role)
}

func TestHTTPBatchUpdateUsers(t *testing.T) {

	data := "[{\"email\":\"example@abtasty.com\",\"role\":\"ADMIN\"},{\"email\":\"example1@abtasty.com\",\"role\":\"VIEWER\"}]"

	_, err := userRequester.HTTPBatchUpdateUsers(data)

	assert.Nil(t, err)
}

func TestHTTPDeleteUser(t *testing.T) {

	err := userRequester.HTTPDeleteUsers("example@abtasty.com")

	assert.Nil(t, err)
}
