package login

import (
	ctx "base/internal/context"
	"github.com/go-chi/chi/v5"
)

type Login struct {
	Context *ctx.Context
}

func NewLogin(context *ctx.Context) *Login {
	return &Login{Context: context}
}

func (l *Login) SetupRouter(router chi.Router) {
	router.Get("/login", l.HandleLoginPage)
	router.Post("/login", l.HandleAuthentication)
}
