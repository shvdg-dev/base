package login

import (
	doc "base/internal/document"
	"net/http"
)

func (l *Login) HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Login"))
	l.context.Renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
}

func (l *Login) HandleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := l.context.UserService.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		http.Redirect(writer, request, "/home", 303)
	} else {
		info := doc.NewInfo(request, doc.WithErrors([]string{"Invalid email or password"}))
		l.context.Renderer.Render(info, l.NewPage(info).CreateLoginPage(), writer, request)
	}
}
