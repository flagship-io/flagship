package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function/web_experimentation"

	mockfunction_ "github.com/flagship-io/flagship/utils/mock_function"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	common.Init(mockfunction_.Auth)

	mockfunction.APICampaign()
	mockfunction.APIToken()

	m.Run()
}
