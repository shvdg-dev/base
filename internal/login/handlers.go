package login

import (
	doc "base/internal/document/info"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, doc.WithTitle("Login"))
	l.Renderer.Render(writer, request, info, l.Views.Login.CreateLoginPage(info))
}

func (l *Login) HandleLoggingIn(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.Context.Users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		l.Context.Sessions.Store("isAuthenticated", true, request)
		info := l.Context.Informer.NewInfo(request)
		l.Renderer.Render(writer, request, info,
			l.Views.Home.CreateHomePage(),
			l.Views.Navbar.CreateInOutButton(info))
	} else {
		info := l.Context.Informer.NewInfo(request, doc.WithErrors([]string{"Invalid email or password"}))
		l.Renderer.Render(writer, request, info, l.Views.Login.CreateLoginPage(info))
	}
}

func (l *Login) HandleLoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store("isAuthenticated", false, request)
	info := l.Context.Informer.NewInfo(request, doc.WithTitle("401 - Unauthorized, whoops!"))
	l.Renderer.Render(writer, request, info,
		l.Views.Error.CreateAuthenticationRequiredPage(),
		l.Views.Navbar.CreateInOutButton(info))
}
