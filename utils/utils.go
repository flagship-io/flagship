package utils

import (
	"encoding/json"

	"github.com/Chadiii/flagship/models"
)

func ListProjectFormat(body []byte) ([]models.Project, error) {
	projectsModel := models.ProjectItems{}

	err := json.Unmarshal(body, &projectsModel)

	if err != nil {
		return nil, err
	}
	return projectsModel, nil
}
