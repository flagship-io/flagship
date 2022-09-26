package models

type Campaign struct {
	ID              string           `json:"id"`
	ProjectID       string           `json:"project_id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Type            string           `json:"type"`
	Status          string           `json:"status"`
	VariationGroups []VariationGroup `json:"variation_groups"`
	Scheduler       Scheduler        `json:"scheduler"`
}

type Scheduler struct {
	StartDate string `json:"start_date"`
	StopDate  string `json:"stop_date"`
	TimeZone  string `json:"timezone"`
}

type CampaignToggleRequest struct {
	State string `json:"state"`
}
