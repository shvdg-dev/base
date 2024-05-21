package login

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Login"))
	rend.GetRenderer().Render(writer, request, info, l.Views.Login.CreateLoginPage(info))
}

func (l *Login) HandleLoggingIn(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.Context.Users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		l.Context.Sessions.Store("isAuthenticated", true, request)
		info := doc.NewInfo(request)
		rend.GetRenderer().Render(writer, request, info, l.Views.Home.CreateHomePage(), l.Views.Navbar.CreateInOutButton(request))
	} else {
		info := doc.NewInfo(request, doc.WithErrors([]string{"Invalid email or password"}))
		rend.GetRenderer().Render(writer, request, info, l.Views.Login.CreateLoginPage(info))
	}
}

func (l *Login) HandleLoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store("isAuthenticated", false, request)
	info := doc.NewInfo(request, doc.WithTitle("401 - Unauthorized, whoops!"))
	rend.GetRenderer().Render(writer, request, info, l.Views.Error.CreateAuthenticationRequiredPage(), l.Views.Navbar.CreateInOutButton(request))
}
