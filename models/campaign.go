package models

type Campaign struct {
}

type CampaignResponse struct {
	Items       []Project `json:"items"`
	CurrentPage int       `json:"current_page"`
}

type CampaignRequest struct {
	Data string `json:"data"`
}

type CampaignToggleRequest struct {
	State string `json:"state"`
}
