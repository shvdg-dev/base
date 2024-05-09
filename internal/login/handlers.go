package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"log"
	"net/http"
)

func HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.Page{
		Path:  "/login",
		Title: "Login",
	}, View(), writer, request)
}

func HandleAuthentication(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")
	log.Println(email, password)
}
