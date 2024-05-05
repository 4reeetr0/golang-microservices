package models

import "time"

type ForecastResponse struct {
	Country             string    `json:"country"`
	City                string    `json:"city"`
	WheatherMain        string    `json:"wheather_main"`
	WheatherDescription string    `json:"wheather_description"`
	Temperature         float64   `json:"temperature"`
	TemperatureMin      float64   `json:"temperature_min"`
	TemperatureMax      float64   `json:"temperature_max"`
	FeelsLike           float64   `json:"feels_like"`
	Humidity            int       `json:"humidity"`
	Pressure            float64   `json:"pressure"`
	Visibility          int       `json:"visibility"`
	Cloudiness          int       `json:"cloudiness"`
	WindSpeed           float64   `json:"wind_speed"`
	WindDeg             float64   `json:"wind_deg"`
	Rain                float64   `json:"rain"`
	Date                time.Time `json:"date"`
}
