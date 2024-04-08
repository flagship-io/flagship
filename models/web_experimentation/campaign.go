package web_experimentation

type CampaignWE struct {
	Id                 int            `json:"id,omitempty"`
	Name               string         `json:"name"`
	Url                string         `json:"url"`
	Description        string         `json:"description"`
	Type               string         `json:"type"`
	SubType            string         `json:"sub_type"`
	State              string         `json:"state"`
	Traffic            *Traffic       `json:"traffic"`
	Variations         *[]VariationWE `json:"variations"`
	SubTests           *[]CampaignWE  `json:"sub_tests"`
	CreatingDate       DateTemplate   `json:"created_at"`
	Labels             []string       `json:"labels"`
	LastPlayTimestamp  DateTemplate   `json:"last_play"`
	LastPauseTimestamp DateTemplate   `json:"last_pause"`
	GlobalCodeCampaign string         `json:"global_code"`
	SourceCode         string         `json:"source_code"`
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
