package handlers

import (
	"base/internal/constants"
	inf "base/internal/document/info"
	data "base/internal/login/data"
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
	title := l.Context.Localizer.Localize(constants.BundleHome)
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(title), inf.WithPath(constants.PathHome))
	homePage := l.Views.Home.CreateHomePage()
	inOutButton := l.Views.Navbar.CreateInOutButton(info)
	l.Renderer.Render(writer, request, info, homePage, inOutButton)
}

// redirectUnauthenticatedUser redirects the unauthenticated user to the login page with the provided email and password.
func (l *Login) redirectUnauthenticatedUser(writer http.ResponseWriter, request *http.Request, email, password string) {
	errorMessage := l.Context.Localizer.Localize(constants.BundleInvalidEmailOrPassword)
	info := l.Context.Informer.NewInfo(request, inf.WithErrors([]string{errorMessage}))
	pageData := data.NewLoginData(data.WithEmail(email), data.WithPassword(password))
	loginPage := l.Views.Login.CreateLoginPage(info, pageData)
	l.Renderer.Render(writer, request, info, loginPage)
}
