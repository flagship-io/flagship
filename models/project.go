package models

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectRequest struct {
	Name string `json:"name"`
}

type ProjectToggleRequest struct {
	State string `json:"state"`
}
