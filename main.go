package main

import (
	"my-go-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	app.Post("/users", handlers.CreateUser)

	app.Listen(":3000")
}
