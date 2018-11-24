package adaptor

import (
	"encoding/json"
	"github.com/mkadiri/golang-microservice/model"
	"net/http"
)

func RequestToUsers(r *http.Request) (users []model.User, err error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&users)

	return users, err
}