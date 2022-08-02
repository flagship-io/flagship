package models

type VariationGroup struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
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
