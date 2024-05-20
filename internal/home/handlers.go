package home

import (
	doc "base/internal/document"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	h.Context.Renderer.Render(
		doc.NewInfo(request, doc.WithTitle("Home")), h.Page(),
		writer, request)
}
