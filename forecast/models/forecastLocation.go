package models

type ForecastLocation struct {
	Id        int     `json:"id"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Lattitude float64 `json:"lattitude"`
	Longitude float64 `json:"longitude"`
}
