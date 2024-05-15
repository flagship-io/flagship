package web_experimentation

type GlobalCode struct {
	Id        int          `json:"id,omitempty"`
	Code      string       `json:"code"`
	CreatedAt DateTemplate `json:"created_at"`
}

type GlobalCodeStr struct {
	GlobalCode string `json:"global_code"`
}

type ModificationCodeStr struct {
	InputType string `json:"input_type"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Selector  string `json:"selector,omitempty"`
	Type      string `json:"type,omitempty"`
	Engine    string `json:"engine,omitempty"`
}
