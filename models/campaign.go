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

type VariationGroup struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Variations []Variation `json:"variations"`
	Targeting  Targeting   `json:"targeting"`
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

type Targeting struct {
	TargetingGroups []TargetingGroup `json:"targeting_groups"`
}

type TargetingGroup struct {
	Targetings []InnerTargeting `json:"targetings"`
}

type InnerTargeting struct {
	Key      string      `json:"key"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type CampaignToggleRequest struct {
	State string `json:"state"`
}
