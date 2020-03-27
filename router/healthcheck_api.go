package router

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := RestResponse{Message: "OK", StatusCode: http.StatusOK}
	response.Write(w)
}