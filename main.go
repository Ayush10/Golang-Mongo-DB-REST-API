package main

import (
	"Golang_Mongo_DB_REST_API/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserControllers(getSession())
	r.GET("/user/id", uc.GetUser)
	r.POST("/user".uc.CreateUser)
	r.DELETE("/user/id".uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}

	return s
}
