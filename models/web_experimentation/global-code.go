package web_experimentation

type GlobalCode struct {
	Id        int          `json:"id,omitempty"`
	Code      string       `json:"code"`
	CreatedAt DateTemplate `json:"created_at"`
}
