package models

type ForecastRequest struct {
	Lattitude float64 `json:"lattitude"`
	Longitude float64 `json:"longitude"`
	Units     string  `json:"units"`
	Lang      string  `json:"lang"`
}
