package httprequest

import (
	"testing"

	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIProject()
	mockfunction.APICampaign()
	mockfunction.APIFlag()
	mockfunction.APITargetingKey()
	mockfunction.APIVariationGroup()
	mockfunction.APIVariation()
	mockfunction.APIToken()
	mockfunction.APIUser()
	mockfunction.APIPanic()
	m.Run()
}
