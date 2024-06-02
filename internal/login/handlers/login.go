package handlers

import (
	"base/internal/constants"
	inf "base/internal/document/info"
	"base/pkg/utils"
	"net/http"
)

// LoggingIn handles the login request by validating the user and redirecting accordingly.
func (l *Login) LoggingIn(writer http.ResponseWriter, request *http.Request) {
	email, password := l.extractCredentials(request)
	if l.isValidUser(email, password) {
		l.redirectAuthenticatedUser(writer, request)
	} else {
		l.redirectUnauthenticatedUser(writer, request, email, password)
	}
}

// extractCredentials extracts the email and password from the request form values.
func (l *Login) extractCredentials(request *http.Request) (string, string) {
	return request.FormValue(constants.ValueEmail), request.FormValue(constants.ValuePassword)
}

// isValidUser validates if the provided email and password are valid for the user.
func (l *Login) isValidUser(email, password string) bool {
	return utils.IsValidEmail(email) && l.Context.Users.IsPasswordCorrect(email, password)
}

// redirectAuthenticatedUser redirects the authenticated user to the home page after successful login.
func (l *Login) redirectAuthenticatedUser(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store(constants.ValueIsAuthenticated, true, request)
	info := l.Context.Informer.NewInfo(request,
		inf.WithTitle(l.Context.Localizer.Localize(constants.BundleHome)),
		inf.WithPath(constants.PathHome))
	l.Renderer.Render(writer, request, info,
		l.Views.Home.CreateHomePage(),
		l.Views.Navbar.CreateInOutButton(info))
}

// redirectUnauthenticatedUser redirects the unauthenticated user to the login page with the provided email and password.
func (l *Login) redirectUnauthenticatedUser(writer http.ResponseWriter, request *http.Request, email, password string) {
	info := l.Context.Informer.NewInfo(request, inf.WithErrors([]string{l.Context.Localizer.Localize(constants.BundleInvalidEmailOrPassword)}))
	data := l.Views.Login.NewData(email, password)
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}
