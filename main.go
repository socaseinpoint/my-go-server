package main

import (
	"my-go-server/database"
	"my-go-server/models"
	"my-go-server/routes"

	_ "my-go-server/docs" // Генерируемая документация

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger" // Swagger middleware
)

// @title My Go Server API
// @version 1.0
// @description Swagger API documentation for My Go Server
// @host localhost:3000
// @BasePath /

func main() {
	app := fiber.New()

	// Middleware для логирования
	app.Use(logger.New())

	// Middleware для CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000", // Указываем конкретные источники
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true, // Поддержка credentials
	}))

	// Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Подключение к базе данных
	database.Connect()

	// Автоматическая миграция модели User
	database.DB.AutoMigrate(&models.User{})

	// Настройка маршрутов
	routes.SetupUserRoutes(app)

	// Запуск приложения
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
