package routes

import (
	"gunmurat7/todo-app-server/controllers"
	"gunmurat7/todo-app-server/repository"
	"gunmurat7/todo-app-server/service"

	"github.com/gofiber/fiber/v2"
)

// Repository -> Service -> Controller

var todocontroller = controllers.NewTodoController(service.NewTodoService(repository.NewTodoRepository("todos")))

func TodoRoutes(app *fiber.App) {
	app.Get("/todos", todocontroller.GetTodos).Get("/todos/:id", todocontroller.GetTodo).Post("/todos", todocontroller.CreateTodo).Put("/todos/:id", todocontroller.UpdateTodo).Delete("/todos/:id", todocontroller.DeleteTodo).Put("/todos/:id/completed", todocontroller.UpdateTodoCompleted)

}
