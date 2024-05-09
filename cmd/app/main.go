package main

import (
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const (
	port = "127.0.0.1:3000"
)

func main() {
	err := http.ListenAndServe(port, setupRouter())
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() chi.Router {
	router := chi.NewRouter()
	files.Router(router)
	home.Router(router)
	docs.Router(router)
	login.Router(router)
	return router
}
