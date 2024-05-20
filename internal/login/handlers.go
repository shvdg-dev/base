package login

import (
	doc "base/internal/document"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Login"))
	l.Context.Renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
}

func (l *Login) HandleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.Context.Users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		l.Context.Sessions.Store("isAuthenticated", true, request)
		http.Redirect(writer, request, "/home", 303)
	} else {
		info := doc.NewInfo(request, doc.WithErrors([]string{"Invalid email or password"}))
		l.Context.Renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
	}
}
