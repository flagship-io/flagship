package web_experimentation

import (
	"net/http"

	models_ "github.com/flagship-io/flagship/models"
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/jarcoal/httpmock"
)

var TestAccount = models_.AccountJSON{
	CurrentUsedCredential: "test_auth",
	AccountID:             "account_id",
	AccountEnvironmentID:  "account_environment_id",
}

var TestGlobalCode = models.GlobalCode_{
	OnDomReady: true,
	Value:      "console.log(\"test\")",
}

var accountID = "account_id"

var TestAccountGlobalCode = models.AccountWE{
	Id:         100000,
	Name:       "account_name",
	Identifier: "account_identifier",
	Role:       "account_role",
	GlobalCode: TestGlobalCode,
}

func APIAccount() {

	httpmock.RegisterResponder("GET", utils.GetWebExperimentationHost()+"/v1/accounts/"+accountID,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestAccountGlobalCode)
			return resp, nil
		},
	)
}
