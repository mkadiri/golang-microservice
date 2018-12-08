package api

import (
	"github.com/mkadiri/golang-microservice/model"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := model.RestResponse{Message: "OK", StatusCode: http.StatusOK}
	response.Write(w)
}