package routes

import (
	"my-go-server/database"
	"my-go-server/models"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/users")

	user.Get("/", func(c *fiber.Ctx) error {
		var users []models.User
		database.DB.Find(&users)
		return c.JSON(users)
	})

	user.Post("/", func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		database.DB.Create(&user)
		return c.JSON(user)
	})

	user.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		database.DB.Delete(&models.User{}, id)
		return c.SendString("User deleted")
	})
}
