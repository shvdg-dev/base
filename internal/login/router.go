package login

import (
	ctx "base/internal/context"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Login struct {
	Context *ctx.Context
	Views   *views.Views
}

func NewLogin(context *ctx.Context, views *views.Views) *Login {
	return &Login{Context: context, Views: views}
}

func (l *Login) SetupRouter(router chi.Router) {
	router.Get("/login", l.HandleLoginPage)
	router.Post("/login", l.HandleLoggingIn)
	router.Get("/logout", l.HandleLoggingOut)
}
