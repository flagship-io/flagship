package models

type Flag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Source      string `json:"source"`
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
