package login

import (
	"base/internal/handlers"
	"base/internal/models"
	"base/internal/views/login"
	"net/http"
)

func Navigation(writer http.ResponseWriter, request *http.Request) {
	handlers.Render(models.Page{
		Path:  "/login",
		Title: "Login",
	}, login.Page(), writer, request)
}
