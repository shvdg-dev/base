package handlers

import (
	"base/internal/constants"
	inf "base/internal/document/info"
	viewData "base/internal/login/data"
	"net/http"
)

// LoginPage handles the login page request.
func (l *Login) LoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(constants.BundleLogin)))
	data := viewData.NewLogin()
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}
