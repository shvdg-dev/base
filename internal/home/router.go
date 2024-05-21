package home

import (
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Home struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

func NewHome(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Home {
	return &Home{Context: context, Views: views, Renderer: renderer}
}

func (h *Home) SetupRouter(router chi.Router) {
	router.Get("/", h.HandleHomePage)
	router.Get("/home", h.HandleHomePage)
}
