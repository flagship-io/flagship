package feature_experimentation

type VariationFE struct {
	Id            string       `json:"id,omitempty"`
	Name          string       `json:"name"`
	Reference     bool         `json:"reference"`
	Allocation    int          `json:"allocation"`
	Modifications Modification `json:"modifications"`
}

type Modification struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
