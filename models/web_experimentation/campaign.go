package web_experimentation

type Test struct {
	Id          int              `json:"id,omitempty"`
	Name        string           `json:"name"`
	Url         string           `json:"url"`
	Description string           `json:"description"`
	Type        string           `json:"type"`
	SubType     string           `json:"sub_type"`
	State       string           `json:"state"`
	Traffic     *Traffic         `json:"traffic"`
	Variations  *[]TestVariation `json:"variations"`
}

type Traffic struct {
	Value                int    `json:"value"`
	LastIncreasedTraffic string `json:"last_increased_traffic"`
	Visitors             int    `json:"visitors"`
	OriginalVisitors     int    `json:"original_visitors"`
	VisitorsLimit        int    `json:"visitors_limit"`
}

type DateTemplate struct {
	ReadableDate string `json:"readable_date"`
	Timestamp    int    `json:"timestamp"`
	Pattern      string `json:"pattern"`
}
