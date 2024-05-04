package docs

import (
	"base/internal/handlers"
	"base/internal/views/docs"
	"net/http"
)

func Navigation(writer http.ResponseWriter, request *http.Request) {
	handlers.Render("Docs", docs.Page(), writer, request)
}
