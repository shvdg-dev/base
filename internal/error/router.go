package error

import (
	ctx "base/internal/context"
	"base/internal/handlers"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

// Error is used for routing and handlers regarding errors.
type Error struct {
	Context  *ctx.Context
	Handlers *handlers.Handlers
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewError returns a new instance of Error.
func NewError(context *ctx.Context, handlers *handlers.Handlers, views *views.Views, renderer *rend.Renderer) *Error {
	return &Error{Context: context, Handlers: handlers, Views: views, Renderer: renderer}
}

// SetupRouter sets up the error router.
func (e *Error) SetupRouter(router chi.Router) {
	router.NotFound(e.Handlers.Error.NotFound())
}
