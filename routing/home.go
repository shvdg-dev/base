package routing

import (
	"base/models"
	"base/views/pages"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	render(models.Page{
		Title:   "Home",
		Content: pages.Home(),
	}, writer, request)
}
