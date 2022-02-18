# Golang-Todo-App
This repo contains a Rest API that is developed with Golang, Go Fiber framework and MongoDB.

## Using

* Clone the repo and create an .env file in the project directory(or where you want but you should give .env file path in the init function as right).

* In env file create an 'MONGO_URI' variable with your mongodb cluster address.

* Install required packages
```
go get
```

* Run service tests
```
go test test/todo_test.go 
```

* Run the project
```
go run server.go
```
