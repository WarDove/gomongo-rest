package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomongo/model"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserController struct {
	Mgoclient *mongo.Client
	Ctx       context.Context
	Cancel    context.CancelFunc
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := model.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = primitive.NewObjectID()

	uc.Mgoclient.Database("gomongo").Collection("users").InsertOne(uc.Ctx, u)
	defer uc.Cancel()
	if err := uc.Mgoclient.Disconnect(uc.Ctx); err != nil {
		panic(err)
	}
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
