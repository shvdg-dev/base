package home

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/", Handler)
	router.Get("/home", Handler)
}
