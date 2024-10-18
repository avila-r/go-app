package main

import (
	"github.com/avila-r/env"
	"github.com/gofiber/fiber/v2"
)

var (
	URL = func() string {
		if err := env.Load(); err != nil {
			panic(err)
		}

		url := env.Get("SERVER_URL")
		if url == "" {
			url = ":8888"
		}

		return url
	}()
)

func main() {
	app := fiber.New()

	app.Get("/greeting", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(URL)
}
