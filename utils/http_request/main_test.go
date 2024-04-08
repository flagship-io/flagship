package http_request

import (
	"testing"

	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	"github.com/flagship-io/flagship/utils/mock_function/web_experimentation"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	feature_experimentation.APIToken()
	web_experimentation.APIToken()
	mockfunction.Request()
	m.Run()
}
