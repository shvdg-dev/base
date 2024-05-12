package login

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/login", handleLoginPage)
	router.Post("/login", handleAuthentication)
}
