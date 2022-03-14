package repository

import (
	"context"
	"encoding/json"
	"errors"
	"gunmurat7/todo-app-server/config"
	"gunmurat7/todo-app-server/helpers"
	"gunmurat7/todo-app-server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetOneTodo(id string) (*models.Todo, error)
	GetAllTodos() ([]models.Todo, error)
	CreateTodo(body []byte) (*models.Todo, error)
	UpdateTodo(id string, body []byte) (*models.Todo, error)
	DeleteTodo(id string) (*models.Todo, error)
}

type TodoRepositoryImpl struct {
	collection string
}

func NewTodoRepository(col string) TodoRepository {
	return &TodoRepositoryImpl{
		collection: col,
	}
}

func (repository *TodoRepositoryImpl) GetOneTodo(id string) (*models.Todo, error) {
	var model models.Todo

	hex, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	data := config.GetCollection(repository.collection).FindOne(context.TODO(), primitive.M{"_id": hex})

	if data.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("Todo not found with id: " + id)
	}

	if data.Err() != nil {
		return nil, data.Err()
	}

	decErr := data.Decode(&model)
	if decErr != nil {
		return nil, errors.New("todo can not decode")
	}

	return &model, nil
}

func (repository *TodoRepositoryImpl) GetAllTodos() ([]models.Todo, error) {
	var modelList []models.Todo

	cursor, err := config.GetCollection(repository.collection).Find(context.TODO(), helpers.SFilter, helpers.Opt)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var model models.Todo
		err := cursor.Decode(&model)
		if err != nil {
			return nil, err
		}
		modelList = append(modelList, model)
	}

	return modelList, nil
}

func (repository *TodoRepositoryImpl) CreateTodo(body []byte) (*models.Todo, error) {
	var model models.Todo

	decErr := json.Unmarshal(body, &model)
	if decErr != nil {
		return nil, errors.New("todo can not decode")
	}

	model.CreatedAt = time.Now()

	res, err := config.GetCollection(repository.collection).InsertOne(context.TODO(), model)
	if err != nil {
		return nil, errors.New("todo can not create")
	}

	model.ID = res.InsertedID.(primitive.ObjectID)

	return &model, nil
}

func (repository *TodoRepositoryImpl) UpdateTodo(id string, body []byte) (*models.Todo, error) {

	// Get todo model with id and update it
	todo, err := repository.GetOneTodo(id)
	if err != nil {
		return nil, err
	}

	decErr := json.Unmarshal(body, &todo)
	if decErr != nil {
		return nil, errors.New("todo can not decode")
	}

	res, err := config.GetCollection(repository.collection).UpdateOne(context.TODO(), primitive.M{"_id": todo.ID}, bson.M{"$set": todo})
	if err != nil {
		return nil, errors.New("todo can not update")
	}

	if res.ModifiedCount == 0 {
		return nil, errors.New("todo can not update")
	}

	return todo, nil
}

func (repository *TodoRepositoryImpl) DeleteTodo(id string) (*models.Todo, error) {

	hex, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	res, err := config.GetCollection(repository.collection).DeleteOne(context.TODO(), primitive.M{"_id": hex})

	if err != nil {
		return nil, err
	}

	if res.DeletedCount == 0 {
		return nil, errors.New("Todo not found with id: " + id)
	}

	return &models.Todo{}, nil

}
