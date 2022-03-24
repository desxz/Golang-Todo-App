package helpers

import (
	"gunmurat7/todo-app-server/config"
	"log"
)

var IsTestEnv = false

func Initialize() {
	// err := godotenv.Load("/Users/muratgun/Dev/GoProjects/todo-app-server/.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	err := config.ConnectMongo()
	if err != nil {
		log.Fatal("Error connecting to mongo")
	}
}
