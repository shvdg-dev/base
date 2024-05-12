package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"base/internal/users"
	"net/http"
)

func handleLoginPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.NewPage("/login", document.WithTitle("Login")), page(), writer, request)
}

func handleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		http.Redirect(writer, request, "/home", 303)
	} else {
		handleLoginPage(writer, request)
	}
}
