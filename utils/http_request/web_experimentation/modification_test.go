package web_experimentation

import (
	"testing"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	"github.com/stretchr/testify/assert"
)

var modificationRequester = ModificationRequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPListModification(t *testing.T) {

	respBody, err := modificationRequester.HTTPListModification(100000)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, web_experimentation.TestModifications.Data.Modifications, respBody)

}

func TestHTTPGetModification(t *testing.T) {

	respBody, err := modificationRequester.HTTPGetModification(100000, 120003)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []models.Modification{web_experimentation.TestModification.Data.Modifications[0]}, respBody)

}
