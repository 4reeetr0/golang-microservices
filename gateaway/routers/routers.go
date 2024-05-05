package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	config "gateaway-service/config"
	models "gateaway-service/models"
	"io/ioutil"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

const authService string = "http://localhost:3001"
const forecastService string = "http://localhost:3002"

func Init(router *fiber.App) {
	publicGroup := router.Group("/api/auth")
	publicGroup.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		response, err := http.Post(authService+"/login", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var authResponse models.AuthResponse
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

	authGroup := router.Group("/api/forecast")
	authGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: config.JWT_SECRET,
		},
	}))

	authGroup.Get("/now", func(c *fiber.Ctx) error {
		body := c.Body()

		response, err := http.Post(forecastService+"/now", "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		var forecastResponse models.ForecastResponse
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer response.Body.Close()
		err = json.Unmarshal(body, &forecastResponse)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(forecastResponse)
	})

	authGroup.Get("/history", func(c *fiber.Ctx) error {
		var body []byte

		response, err := http.Get(forecastService + "/history")
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		var forecastResponse []models.ForecastHistoryResponse
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer response.Body.Close()
		err = json.Unmarshal(body, &forecastResponse)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(forecastResponse)
	})
}
