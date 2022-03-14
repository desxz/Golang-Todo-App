package routes

import (
	"gunmurat7/todo-app-server/controllers"
	"gunmurat7/todo-app-server/helpers"
	"gunmurat7/todo-app-server/repository"
	"gunmurat7/todo-app-server/service"

	"github.com/gofiber/fiber/v2"
)

// Repository -> Service -> Controller
func GetController() controllers.TodoControllerInterface {
	if helpers.IsTestEnv {
		return controllers.NewTodoController(service.NewTodoService(repository.NewMockRepository()))
	}

	return controllers.NewTodoController(service.NewTodoService(repository.NewTodoRepository("todos")))
}

func TodoRoutes(app *fiber.App) {
	app.Get("/todos", GetController().GetTodos).Get("/todos/:id", GetController().GetTodo).Post("/todos", GetController().CreateTodo).Put("/todos/:id", GetController().UpdateTodo).Delete("/todos/:id", GetController().DeleteTodo).Put("/todos/:id/completed", GetController().UpdateTodoCompleted)

}
