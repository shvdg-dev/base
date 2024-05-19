package home

import (
	doc "base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(
		doc.NewInfo(doc.WithPath("/home"), doc.WithTitle("Home")), h.Page(),
		writer, request)
}
