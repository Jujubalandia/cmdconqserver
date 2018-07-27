package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETItems(t *testing.T) {
	t.Run("returns XPTO's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/items/XPTO", nil)
		response := httptest.NewRecorder()
		ItemServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

}
