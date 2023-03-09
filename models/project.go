package models

type Project struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type ProjectToggleRequest struct {
	State string `json:"state"`
}
