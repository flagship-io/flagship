package web_experimentation

import (
	"net/http"
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestVariation = models.VariationWE{
	Id:          110000,
	Name:        "testVariationName",
	Description: "testVariationDescription",
}

func APIVariation() {

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/variations/"+strconv.Itoa(TestVariation.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariation)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/variations/"+strconv.Itoa(TestVariation.Id),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)

}
