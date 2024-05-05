package routers

import (
	controllers "forecast-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func Init(router *fiber.App) {
	router.Get("/now", func(c *fiber.Ctx) error {
		return controllers.GetCurrentWeather(c)
	})

	router.Get("/history", func(c *fiber.Ctx) error {
		return controllers.GetForecastHistory(c)
	})
}
