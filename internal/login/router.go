package login

import (
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Login struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

func NewLogin(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Login {
	return &Login{Context: context, Views: views, Renderer: renderer}
}

func (l *Login) SetupRouter(router chi.Router) {
	router.Get("/login", l.HandleLoginPage)
	router.Post("/login", l.HandleLoggingIn)
	router.Get("/logout", l.HandleLoggingOut)
}
