package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETItems(t *testing.T) {
	t.Run("returns XPTO's score", func(t *testing.T) {
		request := newGetScoreRequest("XPTO")
		response := httptest.NewRecorder()

		ItemServer(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns XPTO2's score", func(t *testing.T) {
		request := newGetScoreRequest("XPTO")
		response := httptest.NewRecorder()

		ItemServer(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/items/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}
