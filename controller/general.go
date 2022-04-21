package controller

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ctrl struct{}

// MongoSession function connects to mongodb server and creates mongo client (session)
func MongoSession(ctx context.Context) *mongo.Client {

	//defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:<you will never see this part :P >/")) // "mongodb://localhost:27017"
	if err != nil {
		panic(err)
	}

	return client
}
