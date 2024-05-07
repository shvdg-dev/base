package docs

import (
	"base/internal/handlers"
	"base/internal/models"
	"base/internal/views/docs"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	handlers.Render(models.Page{
		Path:  "/docs",
		Title: "Docs",
	}, docs.Page(), writer, request)
}
