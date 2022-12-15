package panic

import (
	"testing"

	"github.com/flagship-io/flagship-cli/utils"
	mockfunction "github.com/flagship-io/flagship-cli/utils/mock_function"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIPanic()
	m.Run()
}

func TestPanicCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(PanicCmd)
	assert.Contains(t, output, "Error: required flag(s) \"status\" not set")
}

func TestPanicStatusCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(PanicCmd, "--status=ac")
	assert.Contains(t, failOutput, "Status can only have 2 values: on or off")

	successOutput, _ := utils.ExecuteCommand(PanicCmd, "--status=off")
	assert.Equal(t, "Panic set to off\n", successOutput)
}
