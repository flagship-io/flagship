package models

type Variation struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Reference     bool         `json:"reference"`
	Allocation    int          `json:"allocation"`
	Modifications Modification `json:"modifications"`
}

type Modification struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
