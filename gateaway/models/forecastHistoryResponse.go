package models

type ForecastHistoryResponse struct {
	Country string             `json:"country"`
	City    string             `json:"city"`
	History []ForecastResponse `json:"history"`
}
