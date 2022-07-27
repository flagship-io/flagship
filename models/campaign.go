package models

type Campaign struct {
	ID              string           `json:"id"`
	ProjectID       string           `json:"project_id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Type            string           `json:"type"`
	Status          string           `json:"status"`
	VariationGroups []VariationGroup `json:"variation_groups"`
}

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

type CampaignToggleRequest struct {
	State string `json:"state"`
}
