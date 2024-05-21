package home

import (
	doc "base/internal/document/info"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Home"))
	h.Renderer.Render(writer, request, info, h.Views.Home.CreateHomePage())
}
