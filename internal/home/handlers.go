package home

import (
	consts "base/internal/constants"
	doc "base/internal/document/info"
	"net/http"
)

// HandleHomePage handles the request for the home page.
func (h *Home) HandleHomePage(writer http.ResponseWriter, request *http.Request) {
	info := h.Context.Informer.NewInfo(request, doc.WithTitle(consts.BundleHome))
	page := h.Views.Home.CreateHomePage()
	h.Renderer.Render(writer, request, info, page)
}
