package main

import (
	"my-go-server/database"
	"my-go-server/models"
	"my-go-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	routes.SetupUserRoutes(app)

	app.Listen(":3000")
}
