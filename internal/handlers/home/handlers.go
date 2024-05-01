package home

import (
	"base/internal/handlers"
	"base/internal/views/home"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	handlers.Render("Home", home.Home(), writer, request)
}
