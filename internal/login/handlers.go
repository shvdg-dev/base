package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"base/internal/users"
	"errors"
	"net/http"
)

func handleLoginPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.NewPage("/login", document.WithTitle("Login")), NewView().page(), writer, request)
}

func handleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	isCorrectPassword := users.IsPasswordCorrect(email, password)
	if isCorrectPassword {
		http.Redirect(writer, request, "/home", 303)
	} else {
		errs := []error{errors.New("invalid email or password")}
		page := document.NewPage("/login", document.WithTitle("Login"), document.WithErrors(errs))
		view := NewView(WithErrors(errs))
		renderer.Render(page, view.page(), writer, request)
	}
}
