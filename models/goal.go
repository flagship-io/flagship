package models

type Goal struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
