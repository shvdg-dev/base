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

// HandleLoggingIn handles the login request by validating the user and redirecting accordingly.
func (l *Login) HandleLoggingIn(writer http.ResponseWriter, request *http.Request) {
	email, password := l.extractCredentials(request)
	if l.isValidUser(email, password) {
		l.redirectAuthenticatedUser(writer, request)
	} else {
		l.redirectUnauthenticatedUser(writer, request, email, password)
	}
}

// extractCredentials extracts the email and password from the request form values.
func (l *Login) extractCredentials(request *http.Request) (string, string) {
	return request.FormValue(consts.ValueEmail), request.FormValue(consts.ValuePassword)
}

// isValidUser validates if the provided email and password are valid for the user.
func (l *Login) isValidUser(email, password string) bool {
	return utils.IsValidEmail(email) && l.Context.Users.IsPasswordCorrect(email, password)
}

// redirectAuthenticatedUser redirects the authenticated user to the home page after successful login.
func (l *Login) redirectAuthenticatedUser(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store(consts.ValueIsAuthenticated, true, request)
	info := l.Context.Informer.NewInfo(request,
		inf.WithTitle(l.Context.Localizer.Localize(consts.BundleHome)),
		inf.WithPath(consts.PathHome))
	l.Renderer.Render(writer, request, info,
		l.Views.Home.CreateHomePage(),
		l.Views.Navbar.CreateInOutButton(info))
}

// redirectUnauthenticatedUser redirects the unauthenticated user to the login page with the provided email and password.
func (l *Login) redirectUnauthenticatedUser(writer http.ResponseWriter, request *http.Request, email, password string) {
	info := l.Context.Informer.NewInfo(request, inf.WithErrors([]string{l.Context.Localizer.Localize(consts.BundleInvalidEmailOrPassword)}))
	data := l.Views.Login.NewData(email, password)
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}

// HandleLoggingOut handles the process of logging out a user.
func (l *Login) HandleLoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store(consts.ValueIsAuthenticated, false, request)
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(consts.BundleNotAuthenticatedTitle)), inf.WithPath(consts.PathHome))
	l.Renderer.Render(writer, request, info,
		l.Views.Error.CreateNotAuthenticatedPage(),
		l.Views.Navbar.CreateInOutButton(info))
}
