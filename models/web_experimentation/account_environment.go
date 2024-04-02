package web_experimentation

type AccountEnvironmentWE struct {
	Id               string `json:"id,omitempty"`
	Environment      string `json:"environment"`
	IsMain           bool   `json:"is_main"`
	Panic            bool   `json:"panic"`
	SingleAssignment bool   `json:"single_assignment"`
}
