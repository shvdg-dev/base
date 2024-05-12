package home

import (
	"base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.Page{
		Path:  "/home",
		Title: "Home",
	}, page(), writer, request)
}
