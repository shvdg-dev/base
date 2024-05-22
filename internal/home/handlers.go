package home

import (
	doc "base/internal/document/info"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	info := h.Context.Informer.NewInfo(request, doc.WithTitle("Home"))
	h.Renderer.Render(writer, request, info, h.Views.Home.CreateHomePage())
}
