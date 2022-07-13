package models

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	Items []Project `json:"items"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}
