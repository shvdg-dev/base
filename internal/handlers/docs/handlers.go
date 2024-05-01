package docs

import (
	"base/internal/handlers"
	"base/internal/views/docs"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	handlers.Render("Docs", docs.Docs(), writer, request)
}
