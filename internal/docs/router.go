package docs

import (
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Docs struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

func NewDocs(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Docs {
	return &Docs{Context: context, Views: views, Renderer: renderer}
}

func (d *Docs) SetupRouter(router chi.Router) {
	router.Get("/docs", d.HandleDocsPage)
}
