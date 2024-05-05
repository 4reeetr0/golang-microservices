package controllers

import (
	config "forecast-service/config"
	database "forecast-service/database"
	models "forecast-service/models"
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/gofiber/fiber/v2"
)

func GetCurrentWeather(c *fiber.Ctx) error {
	var curWeatherReq models.ForecastRequest

	if err := c.BodyParser(&curWeatherReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	res, err := getForecastFromRequest(&curWeatherReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func getForecastFromRequest(req *models.ForecastRequest) (models.ForecastResponse, error) {
	var wheatherResponse models.ForecastResponse

	w, err := owm.NewCurrent(req.Units, req.Lang, config.API_KEY)
	if err != nil {
		return wheatherResponse, err
	}

	w.CurrentByCoordinates(&owm.Coordinates{Latitude: req.Lattitude, Longitude: req.Longitude})

	wheatherResponse = models.ForecastResponse{
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
		Date:                time.Unix(int64(w.Dt), 0),
	}

	return wheatherResponse, nil
}

func GetForecastHistory(c *fiber.Ctx) error {
	res, err := database.GetForecastHistory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
