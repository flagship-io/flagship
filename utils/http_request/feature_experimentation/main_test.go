package feature_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"

	mockfunction_ "github.com/flagship-io/flagship/utils/mock_function"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	common.Init(mockfunction_.Auth)

	mockfunction.APIProject()
	mockfunction.APICampaign()
	mockfunction.APIFlag()
	mockfunction.APIGoal()
	mockfunction.APITargetingKey()
	mockfunction.APIVariationGroup()
	mockfunction.APIVariation()
	mockfunction.APIUser()
	mockfunction.APIPanic()
	mockfunction.APIToken()

	m.Run()
}
