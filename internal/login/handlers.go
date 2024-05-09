package login

import (
	"base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.Page{
		Path:  "/login",
		Title: "Login",
	}, View(), writer, request)
}
