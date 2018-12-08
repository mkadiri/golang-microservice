package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var Router *mux.Router

func init()  {
	Router = mux.NewRouter()
	Router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	Router.HandleFunc("/users", getUsers).Methods("GET")
	Router.HandleFunc("/users", addUsers).Methods("PUT")
}

func ExecuteTestRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func CheckTestResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}