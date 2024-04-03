package web_experimentation

type AccountWE struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Role       string `json:"role"`
}
