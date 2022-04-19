package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomongo/model"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserController ctrl

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mgoclient := MongoSession(ctx)

	u := model.User{}
	json.NewDecoder(r.Body).Decode(&u)

	// Check if user already exists
	occurence, err := mgoclient.Database("gomongo1").Collection("users").CountDocuments(ctx, bson.
		D{{"email", "tarlan@huseynov.net"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(occurence)
	if occurence != 0 {

		if err := mgoclient.Disconnect(ctx); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusConflict)

		fmt.Fprintf(w, "%v\n", `{"code":409, "decription":"user exists")`)
		return

	}

	u.Id = primitive.NewObjectID()

	if hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		panic(err)
	} else {
		u.Password = string(hp)
	}

	// Inserting new user
	if _, err := mgoclient.Database("gomongo1").Collection("users").InsertOne(ctx, u); err != nil {
		panic(err)
	}

	if err := mgoclient.Disconnect(ctx); err != nil {
		panic(err)
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mgoclient := MongoSession(ctx)

	if err := mgoclient.Disconnect(ctx); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
