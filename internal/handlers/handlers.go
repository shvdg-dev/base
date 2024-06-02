package handlers

import (
	ctx "base/internal/context"
	. "base/internal/docs/handlers"
	. "base/internal/error/handlers"
	. "base/internal/home/handlers"
	. "base/internal/login/handlers"
	"base/internal/renderer"
	"base/internal/views"
)

// Handlers represents a collection of handlers.
type Handlers struct {
	Home  *Home
	Docs  *Docs
	Login *Login
	Error *Error
}

// NewHandlers creates a new instance of the Handlers structure.
func NewHandlers(context *ctx.Context, views *views.Views, renderer *renderer.Renderer) *Handlers {
	return &Handlers{
		Home:  NewHome(context, views, renderer),
		Login: NewLogin(context, views, renderer),
		Docs:  NewDocs(context, views, renderer),
		Error: NewError(context, views, renderer)}
}
