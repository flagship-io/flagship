package feature_experimentation

import (
	"testing"

	mockfunction "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIProject()
	mockfunction.APICampaign()
	mockfunction.APIFlag()
	mockfunction.APIGoal()
	mockfunction.APITargetingKey()
	mockfunction.APIVariationGroup()
	mockfunction.APIVariation()
	mockfunction.APIUser()
	mockfunction.APIPanic()
	m.Run()
}
