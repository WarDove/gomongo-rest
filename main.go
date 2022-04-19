package main

import (
	"github.com/gomongo/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	uc := &controller.UserController{}

	r := httprouter.New()

	r.POST("/useradd", uc.CreateUser)
	r.GET("/user", uc.GetUser)

	log.Fatal(http.ListenAndServe(":8080", r))

}
