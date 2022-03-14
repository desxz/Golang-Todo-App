package server

import (
	"gunmurat7/todo-app-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartServer(port string) error {

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Authorization, Content-Length, X-Requested-With",
	}))

	routes.TodoRoutes(app)

	err := app.Listen(port)

	if err != nil {
		return err
	}
	return nil
}
