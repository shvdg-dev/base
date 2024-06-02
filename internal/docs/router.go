package docs

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	hand "base/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// Docs is used for routing and handling regarding documentation.
type Docs struct {
	Context  *ctx.Context
	Handlers *hand.Handlers
}

// NewDocs creates a new instance of the Docs struct.
func NewDocs(context *ctx.Context, handlers *hand.Handlers) *Docs {
	return &Docs{Context: context, Handlers: handlers}
}

// SetupRouter sets up the routes for the Docs struct.
func (d *Docs) SetupRouter(router chi.Router) {
	router.Get(consts.PathDocs, d.Handlers.Docs.DocsPage)
}
