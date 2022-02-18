package controllers

import (
	"gunmurat7/todo-app-server/service"
	"gunmurat7/todo-app-server/utils"

	"github.com/gofiber/fiber/v2"
)

type TodoControllerInterface interface {
	GetTodos(ctx *fiber.Ctx) error
	GetTodo(ctx *fiber.Ctx) error
	CreateTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
	UpdateTodoCompleted(ctx *fiber.Ctx) error
}

type TodoController struct {
	todoService service.TodoServiceInterface
}

func NewTodoController(service service.TodoServiceInterface) TodoControllerInterface {
	return &TodoController{
		todoService: service,
	}
}

//@desc		Get all todos
//@route 	GET /todos
//@access 	Public
func (controller *TodoController) GetTodos(ctx *fiber.Ctx) error {
	utils.SearchFilter(ctx.Query("search"))
	page, limit := utils.PaginationWithFiber(ctx, utils.Opt)

	todos, err := controller.todoService.GetTodosMongo()

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err})
	}
	return ctx.Status(200).JSON(fiber.Map{"success": true, "page": page, "limit": limit, "data": todos})
}

//@desc		Get todos
//@route 	GET /todos:id
//@access 	Public
func (controller *TodoController) GetTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todo, err := controller.todoService.GetTodoMongo(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"success": true, "data": todo})
}

//@desc		Create todo
//@route 	POST /todos
//@access 	Public
func (controller *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json", "utf-8")

	createdTodo, err := controller.todoService.CreateTodoMongo(ctx.Body())
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return ctx.Status(201).JSON(fiber.Map{"success": true, "data": createdTodo})
}

//@desc		Update todo
//@route 	PUT /todos:id
//@access 	Public
func (controller *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	todos, err := controller.todoService.UpdateTodoMongo(id, ctx.Body())
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"success": true, "data": todos})
}

//@desc		Delete todo
//@route 	DELETE /todos:id
//@access 	Public
func (controller *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todos, err := controller.todoService.DeleteTodoMongo(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"success": true, "data": todos})
}

func (controller *TodoController) UpdateTodoCompleted(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := controller.todoService.UpdateTodoCompletedMongo(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"success": true, "data": res})
}
