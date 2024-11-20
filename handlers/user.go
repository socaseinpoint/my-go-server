package handlers

import (
	"my-go-server/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	dto := new(models.CreateUserDTO)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := validate.Struct(dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created", "user": dto})
}
