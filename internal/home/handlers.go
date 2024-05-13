package home

import (
	doc "base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func handleHomePage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(
		doc.NewInfo(doc.WithPath("/home"), doc.WithTitle("Home")), Page(),
		writer, request)
}
