package service

import (
	"encoding/json"
	"gunmurat7/todo-app-server/models"
	"gunmurat7/todo-app-server/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type TodoServiceInterface interface {
	GetTodosMongo() ([]models.Todo, error)
	GetTodoMongo(id string) (*models.Todo, error)
	CreateTodoMongo(body []byte) (*models.Todo, error)
	UpdateTodoMongo(id string, body []byte) (*models.Todo, error)
	DeleteTodoMongo(id string) (*models.Todo, error)
	UpdateTodoCompletedMongo(id string) (bool, error)
}

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoServiceInterface {
	return &TodoService{
		repository: repo,
	}
}

func (service *TodoService) GetTodosMongo() ([]models.Todo, error) {
	res, err := service.repository.GetAllTodos()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TodoService) GetTodoMongo(id string) (*models.Todo, error) {
	res, err := service.repository.GetOneTodo(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TodoService) CreateTodoMongo(body []byte) (*models.Todo, error) {
	res, err := service.repository.CreateTodo(body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TodoService) UpdateTodoMongo(id string, body []byte) (*models.Todo, error) {
	res, err := service.repository.UpdateTodo(id, body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TodoService) DeleteTodoMongo(id string) (*models.Todo, error) {
	res, err := service.repository.DeleteTodo(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *TodoService) UpdateTodoCompletedMongo(id string) (bool, error) {
	res, err := service.repository.GetOneTodo(id)
	if err != nil {
		return false, err
	}

	b := bson.M{"completed": !res.Completed}

	upData, upErr := json.Marshal(b)
	if upErr != nil {
		return false, err
	}

	_, uErr := service.repository.UpdateTodo(id, upData)
	if uErr != nil {
		return false, uErr
	}

	return !res.Completed, nil
}
