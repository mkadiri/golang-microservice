package api

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}