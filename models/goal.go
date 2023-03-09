package models

type Goal struct {
	Id       string `json:"id"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
