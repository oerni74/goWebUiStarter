package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewStarterServer()

	log.Fatal(http.ListenAndServe(":8080", server.Mux))
}
