package login

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/login", HandleLoginPage)
	router.Post("/login", HandleAuthentication)
}
