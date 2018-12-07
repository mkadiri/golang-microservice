package controller

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

	response := ExecuteTestRequest(req)

	CheckTestResponseCode(t, http.StatusOK, response.Code)
}