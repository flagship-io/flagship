package feature_experimentation

type FlagAnalyzed struct {
	LineNumber       int    `json:"LineNumber"`
	FlagKey          string `json:"FlagKey"`
	FlagDefaultValue string `json:"FlagDefaultValue"`
	FlagType         string `json:"FlagType"`
	Exists           bool   `json:"Exists"`
}

type FileAnalyzed struct {
	File    string `json:"File"`
	FileURL string `json:"FileURL"`
	Error   error  `json:"Error"`
	Results []FlagAnalyzed
}
