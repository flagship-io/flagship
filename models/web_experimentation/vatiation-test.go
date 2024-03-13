package models

type TestVariation struct {
	Id           int          `json:"id,omitempty"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Type         string       `json:"type"`
	Traffic      int          `json:"traffic"`
	VisualEditor bool         `json:"visual_editor"`
	CodeEditor   bool         `json:"code_editor"`
	Components   *[]Component `json:"components"`
}

type Component struct {
	Id          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Js          string   `json:"js"`
	Css         string   `json:"css"`
	Html        string   `json:"html"`
	Form        string   `json:"form"`
	Options     string   `json:"options"`
}
