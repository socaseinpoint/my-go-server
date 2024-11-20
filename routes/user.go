package routes

import (
	"my-go-server/database"
	"my-go-server/models"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes godoc
// @Summary      Setup User Routes
// @Description  Initializes user-related endpoints
func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/users")

	// Привязываем функции с аннотациями к маршрутам
	user.Get("/", ListUsers)
	user.Post("/", CreateUser)
	user.Delete("/:id", DeleteUser)
}

// ListUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func ListUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Add a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User data"
// @Success      201   {object}  models.User
// @Failure      400   {object}  fiber.Map
// @Router       /users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         users
// @Param        id  path      string  true  "User ID"
// @Success      200  {string}  string  "User deleted"
// @Failure      404  {object}  fiber.Map
// @Router       /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.User{}, id)
	return c.SendString("User deleted")
}
