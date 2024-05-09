package main

import (
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	"log"
	"net/http"
)

const (
	port = "127.0.0.1:3000"
)

func main() {
	setupRouter()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() {
	files.Router()
	home.Router()
	docs.Router()
	login.Router()
}
