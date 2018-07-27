package httpserver

import (
	"fmt"
	"net/http"
)

func ItemServer(w http.ResponseWriter, r *http.Request) {

	item := r.URL.Path[len("/items/"):]
	fmt.Fprint(w, GetItemScore(item))

}

func GetItemScore(item string) string {
	if item == "XPTO" {
		return "20"
	}

	if item == "XPTO2" {
		return "10"
	}

	return ""
}
