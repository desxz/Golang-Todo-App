package main

import (
	"gunmurat7/todo-app-server/helpers"
	"gunmurat7/todo-app-server/server"
	"log"
)

func init() {
	helpers.Initialize()
}

// export PATH=$PATH:/usr/local/go/bin:/Users/muratgun/go/bin

func main() {
	err := server.StartServer(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
