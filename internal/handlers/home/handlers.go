package home

import (
	"base/internal/handlers"
	"base/internal/views/home"
	"net/http"
)

func Navigation(writer http.ResponseWriter, request *http.Request) {
	handlers.Render("Home", home.Page(), writer, request)
}
