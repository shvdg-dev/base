package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func HandleLoginPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.Page{
		Path:  "/login",
		Title: "Login",
	}, View(), writer, request)
}

func HandleAuthentication(writer http.ResponseWriter, request *http.Request) {

}
