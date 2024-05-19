package files

import (
	"github.com/go-chi/chi/v5"
)

// SetupRouter sets up a file server for serving static files.
func SetupRouter(router chi.Router) {
	router.Get("/public/*", Handler)
}
