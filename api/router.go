package api

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()
	Router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	Router.HandleFunc("/users", getUsers).Methods("GET")
	Router.HandleFunc("/users", addUsers).Methods("PUT")
}