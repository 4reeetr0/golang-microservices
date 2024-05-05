package main

import (
	database "forecast-service/database"
	routers "forecast-service/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.Connect()

	go database.AddForecastsHourly()

	routers.Init(app)

	app.Listen(":3002")

	defer database.DB.Close()
}
