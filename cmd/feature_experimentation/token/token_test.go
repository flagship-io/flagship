package token

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	mockfunction_fe "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testToken models.Token

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_fe.APIToken()
	m.Run()
}

func TestTokenCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(TokenCmd)
	assert.Contains(t, output, "Manage your token\n")
}

func TestTokenHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(TokenCmd, "--help")
	assert.Contains(t, output, "Manage your token\n")
}

func TestTokenInfoCommand(t *testing.T) {
	successOutput, _ := utils.ExecuteCommand(TokenCmd, "info")
	err := json.Unmarshal([]byte(successOutput), &testToken)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestToken, testToken)

}
