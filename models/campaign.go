package models

type Campaign struct {
	Id              string            `json:"id,omitempty"`
	ProjectId       string            `json:"project_id"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Type            string            `json:"type"`
	Status          string            `json:"status"`
	VariationGroups *[]VariationGroup `json:"variation_groups"`
	Scheduler       Scheduler         `json:"scheduler"`
}

type Scheduler struct {
	StartDate string `json:"start_date"`
	StopDate  string `json:"stop_date"`
	TimeZone  string `json:"timezone"`
}

type CampaignSwitchRequest struct {
	State string `json:"state"`
}
