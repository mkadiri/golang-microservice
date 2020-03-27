package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiTestHelper struct {}

func (ApiTestHelper) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func (ApiTestHelper) CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected api_response code %d. Got %d\n", expected, actual)
	}
}