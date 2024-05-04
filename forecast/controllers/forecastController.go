package controllers

import (
	models "forecast-service/models"

	owm "github.com/briandowns/openweathermap"
	"github.com/gofiber/fiber/v2"
)

const API_KEY = "b25b375af8d59912cd7edf9a1911cd07"

func GetCurrentWeather(c *fiber.Ctx) error {
	var curWetherReq models.ForecastRequest

	if err := c.BodyParser(&curWetherReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	w, err := owm.NewCurrent(curWetherReq.Units, curWetherReq.Lang, API_KEY)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	w.CurrentByCoordinates(&owm.Coordinates{Latitude: curWetherReq.Lattitude, Longitude: curWetherReq.Longitude})

	wheatherResponse := getInfoFromResponse(w)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Weather fetched successfully",
		"weather": wheatherResponse,
	})
}

func getInfoFromResponse(w *owm.CurrentWeatherData) models.ForecastResponse {
	return models.ForecastResponse{
		Country:             w.Sys.Country,
		City:                w.Name,
		WheatherMain:        w.Weather[0].Main,
		WheatherDescription: w.Weather[0].Description,
		Temperature:         w.Main.Temp,
		TemperatureMin:      w.Main.TempMin,
		TemperatureMax:      w.Main.TempMax,
		FeelsLike:           w.Main.FeelsLike,
		Humidity:            w.Main.Humidity,
		Pressure:            w.Main.Pressure,
		Visibility:          w.Visibility,
		Cloudiness:          w.Clouds.All,
		WindSpeed:           w.Wind.Speed,
		WindDeg:             w.Wind.Deg,
		Rain:                w.Rain.OneH,
	}
}
