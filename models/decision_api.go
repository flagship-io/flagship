package models

type Visitor struct {
	VisitorId      string         `json:"visitor_id"`
	VisitorContext map[string]any `json:"context"`
}

type ModificationInfo struct {
	Type  string         `json:"type"`
	Value map[string]any `json:"value"`
}

type VariationInfo struct {
	Id            string           `json:"id"`
	Modifications ModificationInfo `json:"modifications"`
	Reference     bool             `json:"reference"`
}

type CampaignInfo struct {
	Id               string        `json:"id"`
	Slug             string        `json:"slug"`
	Type             string        `json:"type"`
	VariationGroupId string        `json:"variationGroupId"`
	Variation        VariationInfo `json:"variation"`
}

type DecisionAPIInfo struct {
	VisitorId     string         `json:"visitorId"`
	CampaignInfos []CampaignInfo `json:"campaigns"`
}
