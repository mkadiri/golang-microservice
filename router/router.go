package router

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()
}
