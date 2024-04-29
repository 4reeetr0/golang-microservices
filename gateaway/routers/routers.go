package routers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var authService string = "http://localhost:3001"

type AuthResponse struct {
	Token string `json:"access_token"`
}

func Init(router *fiber.App) {
	publicGroup := router.Group("/api/auth")
	publicGroup.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		response, err := http.Post(authService+"/login", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var authResponse AuthResponse
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer response.Body.Close()
		err = json.Unmarshal(body, &authResponse)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": authResponse.Token,
		})
	})

	publicGroup.Post("/register", func(c *fiber.Ctx) error {
		body := c.Body()

		response, err := http.Post(authService+"/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if response.StatusCode != 200 {
			return c.Status(response.StatusCode).SendString("User already exists")
		}
		return c.Status(fiber.StatusOK).SendString("User registered")
	})

	authGroup := router.Group("/auth")
	authGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte("e1a924cda66b8c86f00da9f9da733051d802253fc903ffb3ea106c71e5d5f6c6"),
		},
	}))

	authGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
