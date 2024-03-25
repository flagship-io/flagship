package feature_experimentation

import (
	"net/http"
	"net/url"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type UserRequester struct {
	*common.ResourceRequest
}

func (u *UserRequester) HTTPListUsers() ([]models.User, error) {
	return common.HTTPGetAllPages[models.User](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + u.AccountID + "/account_environments/" + u.AccountEnvID + "/users")
}

func (u *UserRequester) HTTPBatchUpdateUsers(data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPut, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+u.AccountID+"/account_environments/"+u.AccountEnvID+"/users", []byte(data))
}

func (u *UserRequester) HTTPDeleteUsers(email string) error {
	_, err := common.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+u.AccountID+"/account_environments/"+u.AccountEnvID+"/users?emails[]="+url.QueryEscape(email), nil)
	return err
}
