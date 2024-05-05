package home

import (
	"base/internal/handlers"
	"base/internal/models"
	"base/internal/views/home"
	"net/http"
)

func Navigation(writer http.ResponseWriter, request *http.Request) {
	handlers.Render(models.Page{
		Path:  "/home",
		Title: "Home",
	}, home.Page(), writer, request)
}
