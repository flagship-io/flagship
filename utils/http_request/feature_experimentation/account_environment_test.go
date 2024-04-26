package feature_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var accountEnvironmentFERequester = AccountEnvironmentFERequester{&common.ResourceRequest{AccountID: "account_id", AccountEnvironmentID: "account_environment_id"}}

func TestHTTPGetAccountEnvironment(t *testing.T) {
	respBody, err := accountEnvironmentFERequester.HTTPGetAccountEnvironment("account_environment_id")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "account_environment_id", respBody.Id)
	assert.Equal(t, "account_environment_name", respBody.Environment)
	assert.Equal(t, true, respBody.IsMain)
	assert.Equal(t, false, respBody.SingleAssignment)
}

func TestHTTPListAccountEnvironment(t *testing.T) {
	respBody, err := accountEnvironmentFERequester.HTTPListAccountEnvironment("account_id")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "account_environment_id", respBody[0].Id)
	assert.Equal(t, "account_environment_name", respBody[0].Environment)
	assert.Equal(t, true, respBody[0].IsMain)
	assert.Equal(t, false, respBody[0].SingleAssignment)

	assert.Equal(t, "account_environment_id1", respBody[1].Id)
	assert.Equal(t, "account_environment_name1", respBody[1].Environment)
	assert.Equal(t, true, respBody[1].IsMain)
	assert.Equal(t, false, respBody[1].SingleAssignment)
}
