package routers

import "github.com/gofiber/fiber/v2"

//authService := "http://localhost:3001"

func Init(router *fiber.App) {
	router.Group("/auth")
	router.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("Register")
	})
	router.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
}
