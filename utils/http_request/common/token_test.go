package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPCheckToken(t *testing.T) {

	respBody, err := HTTPCheckToken("token")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "client_id", respBody.ClientID)
	assert.Equal(t, "account_id", respBody.AccountID)
	assert.Equal(t, 0, respBody.ExpiresIn)
	assert.Equal(t, "*", respBody.Scope)
}

func TestHTTPCreateToken(t *testing.T) {
	respBody, err := HTTPCreateToken("client_id", "client_secret", "client_credentials", "*", 86400)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "access_token", respBody.AccessToken)
	assert.Equal(t, "refresh_token", respBody.RefreshToken)
}
