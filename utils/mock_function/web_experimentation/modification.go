package web_experimentation

import (
	"net/http"
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestModificationsJS = models.Modification{
	Id:          120001,
	Name:        "modification",
	Value:       "console.log(\"test modification\")",
	VariationID: 110000,
	Type:        "customScriptNew",
}

var TestModificationsCSS = models.Modification{
	Id:          120002,
	Name:        "modification",
	Value:       ".id{\"color\": \"black\"}",
	VariationID: 110000,
	Type:        "addCSS",
}

var TestElementModification = models.Modification{
	Id:          120003,
	Name:        "modification",
	Value:       "console.log(\"test modification\")",
	VariationID: 110000,
	Selector:    "document.querySelector()",
	Type:        "customScriptNew",
}

var TestData = models.ModificationWE{
	GlobalModifications: []models.Modification{},
	Modifications:       []models.Modification{TestModificationsJS, TestModificationsCSS},
}

var TestModifData = models.ModificationWE{
	GlobalModifications: []models.Modification{},
	Modifications:       []models.Modification{TestElementModification},
}

var TestModifications = models.ModificationDataWE{
	Data: TestData,
}

var TestModification = models.ModificationDataWE{
	Data: TestModifData,
}

func APIModification() {

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/modifications",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestModifications)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/modifications?ids="+strconv.Itoa(TestElementModification.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestModification)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/modifications/"+strconv.Itoa(TestElementModification.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestElementModification)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/modifications/"+strconv.Itoa(TestModificationsJS.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestModificationsJS)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetWebExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/tests/"+strconv.Itoa(TestCampaign.Id)+"/modifications/"+strconv.Itoa(TestModificationsCSS.Id),
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestModificationsCSS)
			return resp, nil
		},
	)
}
