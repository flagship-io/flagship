package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	"github.com/stretchr/testify/assert"
)

var variationRequester = VariationWERequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPGetVariation(t *testing.T) {

	respBody, err := variationRequester.HTTPGetVariation(100000, 110000)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, web_experimentation.TestVariation, respBody)

}
