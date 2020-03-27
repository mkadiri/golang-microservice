package router

import (
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	response := ApiTestHelper{}.ExecuteRequest(req)
	ApiTestHelper{}.CheckResponseCode(t, http.StatusOK, response.Code)
}