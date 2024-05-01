package home

import (
	"base/internal/routes"
	"base/internal/views/pages"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	routes.Render("Home", pages.Home(), writer, request)
}
