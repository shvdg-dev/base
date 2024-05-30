package error

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Error is used for routing and handlers regarding errors.
type Error struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewError returns a new instance of Error.
func NewError(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Error {
	return &Error{Context: context, Views: views, Renderer: renderer}
}

// SetupRouter sets up the error router.
func (e *Error) SetupRouter(router chi.Router) {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		info := e.Context.Informer.NewInfo(r, doc.WithTitle(e.Context.Localizer.Localize(consts.BundlePageNotFoundTitle)))
		page := e.Views.Error.CreatePageNotFound()
		e.Renderer.Render(w, r, info, page)
	})
}
