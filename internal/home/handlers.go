package home

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Home"))
	rend.GetRenderer().Render(writer, request, info, h.Views.Home.CreateHomePage())
}
