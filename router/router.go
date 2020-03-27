package router

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()
	Router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	Router.HandleFunc("/users", UserApi{}.GetUsers).Methods("GET")
	Router.HandleFunc("/users", UserApi{}.AddUsers).Methods("PUT")
}