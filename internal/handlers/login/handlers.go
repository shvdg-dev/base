package login

import (
	"base/internal/handlers"
	"base/internal/views/login"
	"net/http"
)

func Navigation(writer http.ResponseWriter, request *http.Request) {
	handlers.Render("Login", login.Page(), writer, request)
}
