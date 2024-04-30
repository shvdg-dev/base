package main

import (
	"base/internal/server"
	"log"
	"net/http"
)

const (
	port = "127.0.0.1:3000"
)

func main() {
	server.SetupRouter()
	log.Printf("Starting the server on localhost:%s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
