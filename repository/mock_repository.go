package repository

import (
	"gunmurat7/todo-app-server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRepository struct {
}

func NewMockRepository() TodoRepository {
	return &MockRepository{}
}

func (repository *MockRepository) GetOneTodo(id string) (*models.Todo, error) {

	mockedID, _ := primitive.ObjectIDFromHex(id)

	return &models.Todo{
		ID:          mockedID,
		Title:       "Mocked Title",
		Description: "Mocked Description",
		Completed:   false,
		CreatedAt:   time.Time{},
	}, nil
}

func (repository *MockRepository) GetAllTodos() ([]models.Todo, error) {
	return []models.Todo{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Mocked Title",
			Description: "Mocked Description",
			Completed:   false,
			CreatedAt:   time.Time{},
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Mocked Title 2",
			Description: "Mocked Description 2",
			Completed:   false,
			CreatedAt:   time.Time{},
		},
	}, nil
}

func (repository *MockRepository) CreateTodo(body []byte) (*models.Todo, error) {

	return &models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "Mocked Title",
		Description: "Mocked Description",
		Completed:   false,
		CreatedAt:   time.Time{},
	}, nil
}

func (repository *MockRepository) UpdateTodo(id string, body []byte) (*models.Todo, error) {

	mockedID, _ := primitive.ObjectIDFromHex(id)

	return &models.Todo{
		ID:          mockedID,
		Title:       "Mocked Title",
		Description: "Mocked Description",
		Completed:   false,
		CreatedAt:   time.Time{},
	}, nil
}

func (repository *MockRepository) DeleteTodo(id string) (*models.Todo, error) {

	mockedID, _ := primitive.ObjectIDFromHex(id)

	return &models.Todo{
		ID:          mockedID,
		Title:       "",
		Description: "",
		Completed:   false,
		CreatedAt:   time.Time{},
	}, nil
}
