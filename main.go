package main

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"log"
	"net/http"

	"github.com/destafajri/golang-mongodb/controllers"
)

func main() {
	
	r := httprouter.New()
	usercontrol := controllers.NewUserController(getSession())
	r.GET("/user/:id", usercontrol.GetUser)
	r.POST("/user", usercontrol.CreateUser)
	r.DELETE("/user/:id", usercontrol.RemoveUser)

	http.ListenAndServe("localhost:1000", r)
}

func getSession() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    s, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = s.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    return s
}