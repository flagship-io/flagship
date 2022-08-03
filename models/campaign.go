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

type CampaignToggleRequest struct {
	State string `json:"state"`
}
