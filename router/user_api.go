package router

import (
	"github.com/mkadiri/golang-microservice/hello"
	"net/http"
)

type UserApi struct {}

func (UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := hello.UserRepository{}.GetUsers()

	if err != nil {
		response := RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	response := RestResponse{Message: users, StatusCode: http.StatusOK}
	response.Write(w)
}

func (UserApi) AddUsers(w http.ResponseWriter, r *http.Request) {
	adaptedUsers, err := hello.RequestToUsers(r)

	if err != nil {
		response := RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	var savedUsers []hello.User

	for _, user := range adaptedUsers  {
		savedUser, err := hello.UserRepository{}.AddUser(user)

		if err != nil {
			response := RestResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError}
			response.Write(w)
			return
		}

		savedUsers = append(savedUsers, savedUser)
	}

	response := RestResponse{Message: savedUsers, StatusCode: http.StatusOK}
	response.Write(w)
}