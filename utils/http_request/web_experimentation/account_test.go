package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var accountWERequester = AccountWERequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPListAccount(t *testing.T) {

	respBody, err := accountWERequester.HTTPListAccount()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, 100000, respBody[0].Id)
	assert.Equal(t, "account_name", respBody[0].Name)
	assert.Equal(t, "account_identifier", respBody[0].Identifier)
	assert.Equal(t, "account_role", respBody[0].Role)

}
