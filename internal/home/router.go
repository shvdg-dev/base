package home

import (
	ctx "base/internal/context"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Home struct {
	Context *ctx.Context
	Views   *views.Views
}

func NewHome(context *ctx.Context, views *views.Views) *Home {
	return &Home{Context: context, Views: views}
}

func (h *Home) SetupRouter(router chi.Router) {
	router.Get("/", h.HandleHomePage)
	router.Get("/home", h.HandleHomePage)
}
