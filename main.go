package main

import (
	"log"
	"net/http"

	"github.com/bera/concurhttpserver/httpserver"
)

func main() {
	//"github.com/bera/concurhttpserver/server/httpserver"
	handler := http.HandlerFunc(httpserver.PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen n port 5000 %v", err)
	}
}
