package handlers

import (
	ctx "base/internal/context"
	. "base/internal/docs/handlers"
	. "base/internal/login/handlers"
	"base/internal/renderer"
	"base/internal/views"
)

// Handlers represents a collection of handlers.
type Handlers struct {
	Login *Login
	Docs  *Docs
}

// NewHandlers creates a new instance of the Handlers structure.
func NewHandlers(context *ctx.Context, views *views.Views, renderer *renderer.Renderer) *Handlers {
	return &Handlers{
		Login: NewLogin(context, views, renderer),
		Docs:  NewDocs(context, views, renderer)}
}
