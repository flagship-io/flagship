package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
)

type TestRequest struct {
	Name string `json:"name"`
}

func Request() {
	config.SetViper()

	testRequest := TestRequest{
		Name: "TestName",
	}

	testRequest1 := TestRequest{
		Name: "TestName1",
	}

	testRequestList := []TestRequest{testRequest, testRequest1}

	resp := utils.HTTPListResponse[TestRequest]{
		Items:             testRequestList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", "serverURLGet",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, testRequest)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "serverURLList",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)
}
