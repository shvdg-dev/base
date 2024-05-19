package docs

import (
	ctx "base/internal/context"
	"github.com/go-chi/chi/v5"
)

type Docs struct {
	context *ctx.Context
}

func NewDocs(context *ctx.Context) *Docs {
	return &Docs{context: context}
}

func (d *Docs) SetupRouter(router chi.Router) {
	router.Get("/docs", d.handleDocsPage)
}
