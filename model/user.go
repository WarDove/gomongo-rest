package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Birthdate string             `json:"birthdate" bson:"birthdate"`
	Gender    string             `json:"gender" bson:"gender"`
	password  []byte
}

//{"name":"Tarlan", "email":"tarlan@huseynov.net", "birthdate":30091991", "gender":"male"}
