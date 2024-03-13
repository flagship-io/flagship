package feature_experimentation

type Project struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type ProjectSwitchRequest struct {
	State string `json:"state"`
}
