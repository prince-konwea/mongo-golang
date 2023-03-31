package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prince/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.27017")

	if err != nil {
		panic(err)
	}
	return s
}
