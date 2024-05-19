package login

import (
	doc "base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(doc.WithPath("/login"), doc.WithTitle("Login"))
	renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
}

func (l *Login) HandleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.context.UserService.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		http.Redirect(writer, request, "/home", 303)
	} else {
		info := doc.NewInfo(doc.WithErrors([]string{"Invalid email or password"}))
		renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
	}
}
