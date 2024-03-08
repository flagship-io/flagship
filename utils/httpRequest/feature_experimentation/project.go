package feature_experimentation

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListProject() ([]models.Project, error) {
	return HTTPGetAllPages[models.Project](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/projects")
}

func HTTPGetProject(id string) (models.Project, error) {
	return HTTPGetItem[models.Project](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/projects/" + id)
}

func HTTPCreateProject(data []byte) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects", data)
}

func HTTPEditProject(id string, data []byte) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id, data)
}

func HTTPSwitchProject(id, state string) error {
	projectRequest := models.ProjectSwitchRequest{
		State: state,
	}

	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		return err
	}

	_, err = HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id+"/toggle", projectRequestJSON)
	return err
}

func HTTPDeleteProject(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id, nil)
	return err
}
