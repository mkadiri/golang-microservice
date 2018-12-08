package api

import (
	"github.com/mkadiri/golang-microservice/model"
	"github.com/mkadiri/golang-microservice/repository"
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
	"time"
)

func init() {
	router.Router.HandleFunc("/users", addUser).Methods("PUT")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	dob, _ := time.Parse("2018-08-08", "2018-08-08")
	user := model.User{Id: 1, FirstName: "james", LastName: "smith", DateOfBirth: dob}

	user, _ = repository.AddUser(user)

	response := model.RestResponse{Message:user, StatusCode: http.StatusOK}
	response.Write(w)
}