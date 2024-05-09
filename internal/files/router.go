package files

import (
	"github.com/go-chi/chi/v5"
)

// Router sets up a file server for serving static files.
func Router(router chi.Router) {
	router.Get("/public/*", Handler)
}
