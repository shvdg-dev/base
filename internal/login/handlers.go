package login

import (
	consts "base/internal/constants"
	inf "base/internal/document/info"
	"base/pkg/utils"
	"net/http"
)

// HandleLoginPage handles the login page request.
func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(consts.BundleLogin)))
	data := l.Views.Login.NewData("", "")
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}

// HandleLoggingIn handles the process of logging the user.
func (l *Login) HandleLoggingIn(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue(consts.ValueEmail)
	password := request.FormValue(consts.ValuePassword)
	isValidEmail := utils.IsValidEmail(email)
	isCorrectPassword := l.Context.Users.IsPasswordCorrect(email, password)
	if isValidEmail && isCorrectPassword {
		l.Context.Sessions.Store(consts.ValueIsAuthenticated, true, request)
		info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(consts.BundleHome)), inf.WithPath(consts.PathHome))
		l.Renderer.Render(writer, request, info,
			l.Views.Home.CreateHomePage(),
			l.Views.Navbar.CreateInOutButton(info))
	} else {
		info := l.Context.Informer.NewInfo(request, inf.WithErrors([]string{l.Context.Localizer.Localize(consts.BundleInvalidEmailOrPassword)}))
		data := l.Views.Login.NewData(email, password)
		page := l.Views.Login.CreateLoginPage(info, data)
		l.Renderer.Render(writer, request, info, page)
	}
}

// HandleLoggingOut handles the process of logging out a user.
func (l *Login) HandleLoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store(consts.ValueIsAuthenticated, false, request)
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(consts.BundleNotAuthenticatedTitle)), inf.WithPath(consts.PathHome))
	l.Renderer.Render(writer, request, info,
		l.Views.Error.CreateNotAuthenticatedPage(),
		l.Views.Navbar.CreateInOutButton(info))
}
