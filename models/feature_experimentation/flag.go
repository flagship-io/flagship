package feature_experimentation

type Flag struct {
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Source           string   `json:"source"`
	DefaultValue     string   `json:"default_value,omitempty"`
	PredefinedValues []string `json:"predefined_values,omitempty"`
}

type MultiFlagRequest struct {
	Flags []Flag `json:"flags"`
}

type MultiFlagResponse struct {
	CreatedIds []string `json:"created_ids"`
}
