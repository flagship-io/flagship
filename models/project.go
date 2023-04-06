package models

type Project struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type ProjectToggleRequest struct {
	State string `json:"state"`
}
