package main

import (
	routers "gateaway-service/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routers.Init(app)

	app.Listen(":3000")
}
