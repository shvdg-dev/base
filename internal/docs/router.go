package docs

import (
	"github.com/go-chi/chi/v5"
)

func Router(router chi.Router) {
	router.Get("/docs", handleDocsPage)
}
