package web_experimentation

type AccountWE struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Identifier bool   `json:"identifier"`
	Role       bool   `json:"role"`
}
