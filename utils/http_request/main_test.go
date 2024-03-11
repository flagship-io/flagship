package http_request

import (
	"testing"

	mockfunction "github.com/flagship-io/flagship/utils/mock_function"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIToken()
	mockfunction.Request()
	m.Run()
}
