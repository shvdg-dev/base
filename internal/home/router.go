package home

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

// Home is used for routing and handlers regarding the home page.
type Home struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewHome creates a new instance of the Home.
func NewHome(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Home {
	return &Home{Context: context, Views: views, Renderer: renderer}
}

// SetupRouter sets up the routes for the Home struct.
func (h *Home) SetupRouter(router chi.Router) {
	router.Get("/", h.HandleHomePage)
	router.Get(consts.PathHome, h.HandleHomePage)
}
