package controller

import (
	"bytes"
	"encoding/json"
	"github.com/mkadiri/golang-microservice/model"
	"net/http"
	"testing"
	"time"
)

var dob,_ = time.Parse("2018-08-08", "2018-08-08")

var users = []model.User {
	{Id: 1, FirstName: "james", LastName: "smith", DateOfBirth: dob},
	{Id: 2, FirstName: "bruce", LastName: "wayne", DateOfBirth: dob},
	{Id: 3, FirstName: "jeremy", LastName: "renner", DateOfBirth: dob},
}

func TestGetUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)

	response := ExecuteTestRequest(req)

	CheckTestResponseCode(t, http.StatusOK, response.Code)
}

func TestAddUsers(t *testing.T) {
	modulesJson, err := json.Marshal(users)

	if err != nil {
		t.Errorf("Cannot encode to JSON")
	}

	req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(modulesJson))
	response := ExecuteTestRequest(req)

	CheckTestResponseCode(t, http.StatusOK, response.Code)

	// json.NewEncoder(w).Encode() adds an extra line (\n), we should do the same here when making a comparison
	modulesJsonString := string(modulesJson) + "\n"
	body := response.Body.String()

	if body != modulesJsonString {
		t.Errorf("Expected %s", modulesJsonString)
		t.Errorf("Got %s", body)
	}
}