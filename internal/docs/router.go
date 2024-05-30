package docs

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

// Docs is used for routing and handlers regarding documentation.
type Docs struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewDocs creates a new instance of the Docs struct.
func NewDocs(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Docs {
	return &Docs{Context: context, Views: views, Renderer: renderer}
}

// SetupRouter sets up the routes for the Docs struct.
func (d *Docs) SetupRouter(router chi.Router) {
	router.Get(consts.PathDocs, d.HandleDocsPage)
}
