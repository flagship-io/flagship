package models

type VariationGroup struct {
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Variations []Variation `json:"variations"`
	Targeting  Targeting   `json:"targeting"`
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
