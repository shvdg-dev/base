package handlers

import (
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
)

// Error is used for handlers regarding errors.
type Error struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewError creates a new instance of the Error.
func NewError(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Error {
	return &Error{Context: context, Views: views, Renderer: renderer}
}
