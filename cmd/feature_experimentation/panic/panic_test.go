package panic

import (
	"testing"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	mockfunction_fe "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_fe.APIPanic()
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
