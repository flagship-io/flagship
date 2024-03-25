package feature_experimentation

import (
	"encoding/json"
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type ProjectRequester struct {
	*common.ResourceRequest
}

func (p *ProjectRequester) HTTPListProject() ([]models.Project, error) {
	return common.HTTPGetAllPages[models.Project](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + p.AccountID + "/projects")
}

func (p *ProjectRequester) HTTPGetProject(id string) (models.Project, error) {
	return common.HTTPGetItem[models.Project](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + p.AccountID + "/projects/" + id)
}

func (p *ProjectRequester) HTTPCreateProject(data []byte) ([]byte, error) {
	return common.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+p.AccountID+"/projects", data)
}

func (p *ProjectRequester) HTTPEditProject(id string, data []byte) ([]byte, error) {
	return common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+p.AccountID+"/projects/"+id, data)
}

func (p *ProjectRequester) HTTPSwitchProject(id, state string) error {
	projectRequest := models.ProjectSwitchRequest{
		State: state,
	}

	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		return err
	}

	_, err = common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+p.AccountID+"/projects/"+id+"/toggle", projectRequestJSON)
	return err
}

func (p *ProjectRequester) HTTPDeleteProject(id string) error {
	_, err := common.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+p.AccountID+"/projects/"+id, nil)
	return err
}
