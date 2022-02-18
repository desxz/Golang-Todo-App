package test

import (
	"gunmurat7/todo-app-server/models"
	"gunmurat7/todo-app-server/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockTodoRepository struct {
	mock.Mock
}

func (mock *mockTodoRepository) GetAllTodos() ([]models.Todo, error) {
	args := mock.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (mock *mockTodoRepository) GetOneTodo(id string) (*models.Todo, error) {
	args := mock.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (mock *mockTodoRepository) CreateTodo(body []byte) (*models.Todo, error) {
	args := mock.Called(body)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (mock *mockTodoRepository) UpdateTodo(id string, body []byte) (*models.Todo, error) {
	args := mock.Called(id, body)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (mock *mockTodoRepository) DeleteTodo(id string) (*models.Todo, error) {
	args := mock.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func TestGetTodosMongo(t *testing.T) {

	mockedRepo := new(mockTodoRepository)

	mockedRepo.On("GetAllTodos").Return([]models.Todo{
		{
			ID:          primitive.NewObjectID(),
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "test2",
			Description: "test2",
			Completed:   true,
			CreatedAt:   time.Now(),
		},
	}, nil)

	todoService := service.NewTodoService(
		mockedRepo,
	)

	res, _ := todoService.GetTodosMongo()

	mockedRepo.AssertExpectations(t)

	assert.Equal(t, 2, len(res))

}

func TestGetTodoMongo(t *testing.T) {

	mockedRepo := new(mockTodoRepository)

	mockedRepo.On("GetOneTodo", "5e7a2e0b3f3a9c0017f8f8f4").Return(&models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "test",
		Description: "test",
		Completed:   false,
		CreatedAt:   time.Now(),
	}, nil)

	todoService := service.NewTodoService(
		mockedRepo,
	)

	res, _ := todoService.GetTodoMongo("5e7a2e0b3f3a9c0017f8f8f4")

	mockedRepo.AssertExpectations(t)

	assert.Equal(t, "test", res.Title)

}

func TestCreateTodoMongo(t *testing.T) {

	mockedRepo := new(mockTodoRepository)

	mockedRepo.On("CreateTodo", []byte(`{"title":"test","description":"test"}`)).Return(&models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "test",
		Description: "test",
		Completed:   false,
		CreatedAt:   time.Now(),
	}, nil)

	todoService := service.NewTodoService(
		mockedRepo,
	)

	res, _ := todoService.CreateTodoMongo([]byte(`{"title":"test","description":"test"}`))

	mockedRepo.AssertExpectations(t)

	assert.Equal(t, "test", res.Title)

}

func TestUpdateTodoMongo(t *testing.T) {

	mockedRepo := new(mockTodoRepository)

	id, _ := primitive.ObjectIDFromHex("5e7a2e0b3f3a9c0017f8f8f4")

	mockedRepo.On("UpdateTodo", "5e7a2e0b3f3a9c0017f8f8f4", []byte(`{"title":"test","description":"test"}`)).Return(&models.Todo{
		ID:          id,
		Title:       "test",
		Description: "test",
		Completed:   false,
		CreatedAt:   time.Now(),
	}, nil)

	todoService := service.NewTodoService(
		mockedRepo,
	)

	res, _ := todoService.UpdateTodoMongo("5e7a2e0b3f3a9c0017f8f8f4", []byte(`{"title":"test","description":"test"}`))

	mockedRepo.AssertExpectations(t)

	assert.Equal(t, "test", res.Title)

}

func TestDeleteTodoMongo(t *testing.T) {

	mockedRepo := new(mockTodoRepository)

	id, _ := primitive.ObjectIDFromHex("000000000000000000000000")

	mockedRepo.On("DeleteTodo", "5e7a2e0b3f3a9c0017f8f8f4").Return(&models.Todo{
		ID:          id,
		Title:       "",
		Description: "",
		Completed:   false,
	}, nil)

	todoService := service.NewTodoService(
		mockedRepo,
	)

	res, _ := todoService.DeleteTodoMongo("5e7a2e0b3f3a9c0017f8f8f4")
	mockedRepo.AssertExpectations(t)

	assert.Equal(t, "", res.Title)

}
