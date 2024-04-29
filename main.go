package main

import (
	"log"
	"net/http"
	"tab-collector/routing"
)

const (
	port = "127.0.0.1:3000"
)

func main() {
	routing.SetupRouting()
	log.Printf("Starting the server on localhost:%s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
