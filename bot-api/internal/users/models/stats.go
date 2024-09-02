package models

type StatCard struct {
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Icon     string  `json:"icon"`
	Subtitle string  `json:"subtitle"`
	Type     string  `json:"type"`
}
