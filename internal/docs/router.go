package docs

import (
	ctx "base/internal/context"
	"github.com/go-chi/chi/v5"
)

type Docs struct {
	Context *ctx.Context
}

func NewDocs(context *ctx.Context) *Docs {
	return &Docs{Context: context}
}

func (d *Docs) SetupRouter(router chi.Router) {
	router.Get("/docs", d.HandleDocsPage)
}
