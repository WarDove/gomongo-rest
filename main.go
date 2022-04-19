package main

import (
	"context"
	"github.com/gomongo/controller"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

// MongoSession function connects to mongodb server and creates mongo client (session)
func MongoSession(uri string) (*mongo.Client, context.Context, context.CancelFunc) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri)) // "mongodb://localhost:27017"
	if err != nil {
		panic(err)
	}

	return client, ctx, cancel
}

func main() {
	client, ctx, cancel := MongoSession("mongodb://admin:Westside592@10.10.15.52/")
	uc := &controller.UserController{client, ctx, cancel}
	r := httprouter.New()
	r.POST("/useradd", uc.CreateUser)
	log.Fatal(http.ListenAndServe(":8080", r))

}
