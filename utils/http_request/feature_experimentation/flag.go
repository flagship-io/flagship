package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type FlagRequester struct {
	*common.ResourceRequest
}

func (f *FlagRequester) HTTPListFlag() ([]models.Flag, error) {
	return common.HTTPGetAllPages[models.Flag](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + f.AccountID + "/flags")
}

func (f *FlagRequester) HTTPGetFlag(id string) (models.Flag, error) {
	return common.HTTPGetItem[models.Flag](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + f.AccountID + "/flags/" + id)
}

func (f *FlagRequester) HTTPCreateFlag(data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+f.AccountID+"/flags", []byte(data))
}

func (f *FlagRequester) HTTPEditFlag(id, data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+f.AccountID+"/flags/"+id, []byte(data))
}

func (f *FlagRequester) HTTPDeleteFlag(id string) error {
	_, err := common.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+f.AccountID+"/flags/"+id, nil)
	return err
}
