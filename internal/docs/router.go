package docs

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	hand "base/internal/handlers"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

// Docs is used for routing and hand regarding documentation.
type Docs struct {
	Context  *ctx.Context
	Handlers *hand.Handlers
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewDocs creates a new instance of the Docs struct.
func NewDocs(context *ctx.Context, handlers *hand.Handlers, views *views.Views, renderer *rend.Renderer) *Docs {
	return &Docs{Context: context, Handlers: handlers, Views: views, Renderer: renderer}
}

// SetupRouter sets up the routes for the Docs struct.
func (d *Docs) SetupRouter(router chi.Router) {
	router.Get(consts.PathDocs, d.HandleDocsPage)
}
