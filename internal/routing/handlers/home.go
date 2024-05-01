package handlers

import (
	"base/internal/views/pages"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	render("Home", pages.Home(), writer, request)
}
