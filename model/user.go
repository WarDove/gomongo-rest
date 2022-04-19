package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Birthdate string             `json:"birthdate" bson:"birthdate"`
	Gender    string             `json:"gender" bson:"gender"`
	Admin     bool               `json:"admin" bson:"admin"`
	Password  string             `json:"password" bson:"password"`
}

// curl -X POST -H "Content-Type: application/json" -d '{"name":"Tarlan1", "email":"tarlan@huseynov.net", "birthdate":"30091991", "gender":"male", "admin":true , "password":"Westside592"}' http://localhost:8080/useradd
