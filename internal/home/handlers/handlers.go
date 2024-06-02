package handlers

import (
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
)

// Home is used for handlers regarding the home page.
type Home struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewHome creates a new instance of the Home.
func NewHome(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Home {
	return &Home{Context: context, Views: views, Renderer: renderer}
}
