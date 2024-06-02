package handlers

import (
	ctx "base/internal/context"
	. "base/internal/login/handlers"
	"base/internal/renderer"
	"base/internal/views"
)

// Handlers represents a collection of handlers, made accessible throughout the app.
type Handlers struct {
	Login *Login
}

// NewHandlers creates a new instance of the Handlers structure.
func NewHandlers(context *ctx.Context, views *views.Views, renderer *renderer.Renderer) *Handlers {
	return &Handlers{Login: NewLogin(context, views, renderer)}
}
