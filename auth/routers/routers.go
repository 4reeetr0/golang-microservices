package routers

import (
	controllers "auth-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func Init(router *fiber.App) {
	router.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c)
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})
}
