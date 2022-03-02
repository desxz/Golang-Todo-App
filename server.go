package main

import (
	"gunmurat7/todo-app-server/config"
	"gunmurat7/todo-app-server/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = config.ConnectMongo()
	if err != nil {
		log.Fatal("Error connecting to mongo")
	}

}

// export PATH=$PATH:/usr/local/go/bin:/Users/muratgun/go/bin

func main() {

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Authorization, Content-Length, X-Requested-With",
	}))

	PORT := os.Getenv("PORT")

	routes.TodoRoutes(app)

	err := app.Listen(PORT)

	if err != nil {
		log.Fatal(err)
	}

}
