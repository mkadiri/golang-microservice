package api

import (
	"encoding/json"
	"github.com/mkadiri/golang-microservice/model"
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/users", nil)

	response := executeTestRequest(req)

	dob, _ := time.Parse("2018-08-08", "2018-08-08")
	user := model.User{Id: 1, FirstName: "james", LastName: "smith", DateOfBirth: dob}
	userJson, err := json.Marshal(user)


	if err != nil {
		t.Errorf("Cannot encode to JSON")
	}

	// json.NewEncoder(w).Encode() adds an extra line (\n), we should do the same here when making a comparison
	userJsonString := string(userJson) + "\n"
	body := response.Body.String()

	if body != userJsonString {
		t.Errorf("Expected %s", userJsonString)
		t.Errorf("Got %s", body)
	}
}

func executeTestRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	router.Router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}