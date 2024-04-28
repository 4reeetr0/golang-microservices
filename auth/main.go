package main

import (
	database "auth-service/database"
	routers "auth-service/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.Connect()

	routers.Init(app)

	app.Listen(":3001")

	defer database.DB.Close()
}
