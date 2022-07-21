package models

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	Items       []Project `json:"items"`
	CurrentPage int       `json:"current_page"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}

type ProjectToggleRequest struct {
	State string `json:"state"`
}
