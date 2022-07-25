package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListProject() ([]models.Project, error) {
	return HTTPGetAllPages[models.Project](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/projects")
}

func HTTPGetProject(id string) (models.Project, error) {
	return HTTPGetItem[models.Project](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/projects/" + id)
}

func HTTPCreateProject(name string) error {
	projectRequest := models.ProjectRequest{
		Name: name,
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		return err
	}
	_, err = HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects", projectRequestJSON)
	return err
}

func HTTPEditProject(id, name string) error {
	projectRequest := models.ProjectRequest{
		Name: name,
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		return err
	}
	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id, projectRequestJSON)
	return err
}

func HTTPToggleProject(id, state string) error {
	projectRequest := models.ProjectToggleRequest{
		State: state,
	}

	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		return err
	}

	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id+"/toggle", projectRequestJSON)
	return err
}

func HTTPDeleteProject(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+id, nil)
	return err
}
