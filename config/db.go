package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func ConnectMongo() error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user:user123@cluster0.8keti.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return err

	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err

	}

	fmt.Println("MongoDB Connected Successfully")

	Client = client
	return nil

}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("todo-app").Collection(collectionName)
}
