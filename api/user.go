package api

import (
	"github.com/mkadiri/golang-microservice/adaptor"
	"github.com/mkadiri/golang-microservice/model"
	"github.com/mkadiri/golang-microservice/repository"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetUsers()

	if err != nil {
		response := model.RestResponse{Message:err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	response := model.RestResponse{Message: users, StatusCode: http.StatusOK}
	response.Write(w)
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	adaptedUsers, err := adaptor.RequestToUsers(r)

	if err != nil {
		response := model.RestResponse{Message:err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	var savedUsers []model.User

	for _, user := range adaptedUsers  {
		savedUser, err := repository.AddUser(user)

		if err != nil {
			response := model.RestResponse{Message:err.Error(), StatusCode: http.StatusInternalServerError}
			response.Write(w)
			return
		}

		savedUsers = append(savedUsers, savedUser)
	}

	response := model.RestResponse{Message:savedUsers, StatusCode: http.StatusOK}
	response.Write(w)
}