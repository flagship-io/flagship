package models

type Flag struct {
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Source           string   `json:"source"`
	DefaultValue     string   `json:"default_value,omitempty"`
	PredefinedValues []string `json:"predefined_values,omitempty"`
}

type FlagUsage struct {
	ID                string `json:"id"`
	FlagKey           string `json:"flag_key"`
	Repository        string `json:"repository"`
	FilePath          string `json:"file_path"`
	Branch            string `json:"branch"`
	Line              string `json:"line"`
	CodeLineHighlight string `json:"code_line_highlight"`
	Code              string `json:"code"`
}
