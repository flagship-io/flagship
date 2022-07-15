package models

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	Items       []Project `json:"items"`
	CurrentPage string    `json:"current_page"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}
