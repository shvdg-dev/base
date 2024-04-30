package server

import (
	"base/internal/models"
	"base/internal/views/pages"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	render(models.Component{
		Title:   "Home",
		Content: pages.Home(),
	}, writer, request)
}
