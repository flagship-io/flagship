package token

import (
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var testToken models.Token

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIToken()
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
