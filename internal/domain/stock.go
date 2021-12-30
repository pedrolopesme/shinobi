package domain

type Stock struct {
	Symbol    string `json:"symbol"`
	Company   string `json:"company"`
	Sector    string `json:"sector"`
	Subsector string `json:"subsector"`
	Segment   string `json:"segment"`
}
