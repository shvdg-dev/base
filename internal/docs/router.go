package docs

import (
	ctx "base/internal/context"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

type Docs struct {
	Context *ctx.Context
	Views   *views.Views
}

func NewDocs(context *ctx.Context, views *views.Views) *Docs {
	return &Docs{Context: context, Views: views}
}

func (d *Docs) SetupRouter(router chi.Router) {
	router.Get("/docs", d.HandleDocsPage)
}
