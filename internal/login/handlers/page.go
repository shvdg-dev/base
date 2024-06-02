package handlers

import (
	"base/internal/constants"
	info2 "base/internal/document/info"
	"net/http"
)

// LoginPage handles the login page request.
func (l *Login) LoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, info2.WithTitle(l.Context.Localizer.Localize(constants.BundleLogin)))
	data := l.Views.Login.NewData("", "")
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}
