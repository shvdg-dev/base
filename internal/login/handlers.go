package login

import (
	doc "base/internal/document"
	"base/internal/renderer"
	"base/internal/users"
	"errors"
	"net/http"
)

func handleLoginPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(doc.WithPath("/login"), doc.WithTitle("Login"))
	renderer.Render(info, NewLoginPage(info).CreateLoginPage(), writer, request)
}

func handleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		http.Redirect(writer, request, "/home", 303)
	} else {
		info := doc.NewInfo(doc.WithErrors([]error{errors.New("invalid email or password")}))
		renderer.Render(info, NewLoginPage(info).CreateLoginPage(), writer, request)
	}
}
