package web_experimentation

type GlobalCode struct {
	Id        int          `json:"id,omitempty"`
	Code      string       `json:"code"`
	CreatedAt DateTemplate `json:"created_at"`
}

type GlobalCodeStr struct {
	GlobalCode string `json:"global_code"`
}

type ModificationCodeEditStruct struct {
	InputType string `json:"input_type"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Selector  string `json:"selector,omitempty"`
	Type      string `json:"type,omitempty"`
	Engine    string `json:"engine,omitempty"`
}

type ModificationCodeCreateStruct struct {
	InputType   string `json:"input_type"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Selector    string `json:"selector"`
	Type        string `json:"type"`
	Engine      string `json:"engine"`
	VariationID int    `json:"variation_id"`
}
