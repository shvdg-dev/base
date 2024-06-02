package handlers

import (
	"base/internal/constants"
	info2 "base/internal/info"
	"net/http"
)

// HomePage handles the request for the home page.
func (h *Home) HomePage(writer http.ResponseWriter, request *http.Request) {
	info := h.Context.Informer.NewInfo(request, info2.WithTitle(constants.BundleHome))
	page := h.Views.Home.CreateHomePage()
	h.Renderer.Render(writer, request, info, page)
}
