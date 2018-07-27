package main

import (
	"log"
	"net/http"

	"github.com/bera/concurhttpserver/httpserver"
)

func main() {

	handler := http.HandlerFunc(httpserver.ItemServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen n port 5000 %v", err)
	}
}
