package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"base/internal/users"
	"log"
	"net/http"
)

func handleLoginPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.Page{
		Path:  "/login",
		Title: "Login",
	}, page(), writer, request)
}

func handleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	success := users.AuthenticateUser(email, password)
	log.Println(success)
}
