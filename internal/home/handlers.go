package home

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	rend.GetRenderer().Render(
		doc.NewInfo(request, doc.WithTitle("Home")), h.CreateHomePage(),
		writer, request)
}
