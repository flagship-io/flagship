package feature_experimentation

type TargetingKey struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
