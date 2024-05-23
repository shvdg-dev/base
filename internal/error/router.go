package error

import (
	ctx "base/internal/context"
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Error struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

func NewError(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Error {
	return &Error{Context: context, Views: views, Renderer: renderer}
}

func (l *Error) SetupRouter(router chi.Router) {
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		info := l.Context.Informer.NewInfo(r, doc.WithTitle("404 - Page not found!"))
		l.Renderer.Render(w, r, info, l.Views.Error.CreatePageNotFound())
	})
}
