package feature_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var panicRequester = PanicRequester{&common.ResourceRequest{AccountID: "account_id", AccountEnvironmentID: "account_environment_id"}}

func TestHTTPUpdatePanic(t *testing.T) {

	_, err := panicRequester.HTTPUpdatePanic("active")

	assert.Nil(t, err)
}
