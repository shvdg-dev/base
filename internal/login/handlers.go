package login

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Login"))
	rend.GetRenderer().Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
}

func (l *Login) HandleLoggingIn(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.Context.Users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		l.Context.Sessions.Store("isAuthenticated", true, request)
		http.Redirect(writer, request, "/home", 303)
	} else {
		info := doc.NewInfo(request, doc.WithErrors([]string{"Invalid email or password"}))
		rend.GetRenderer().Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
	}
}

func (l *Login) HandleLoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store("isAuthenticated", false, request)
	http.Redirect(writer, request, "/home", 303)
	return
}
