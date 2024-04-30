package server

import (
	"base/internal/views/pages"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	render("Home", pages.Home(), writer, request)
}
