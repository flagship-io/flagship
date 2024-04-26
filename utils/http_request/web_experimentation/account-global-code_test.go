package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var accountGlobalCodeRequester = AccountGlobalCodeRequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPAccountGlobalCode(t *testing.T) {

	respBody, err := accountGlobalCodeRequester.HTTPGetAccountGlobalCode("account_id")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "console.log(\"test\")", respBody)

}
