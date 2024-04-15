package web_experimentation

import (
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type ModificationRequester struct {
	*common.ResourceRequest
}

func (m *ModificationRequester) HTTPGetModification(id int) ([]models.Modification, error) {
	resp, err := common.HTTPGetItem[models.ModificationDataWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + m.AccountID + "/tests/" + strconv.Itoa(id) + "/modifications")
	return resp.Data.Modifications, err
}
