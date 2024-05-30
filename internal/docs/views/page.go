package views

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Docs is used for views regarding documentation.
type Docs struct {
	Context *ctx.Context
}

// NewDocs creates a new Docs instance with the given context.
func NewDocs(context *ctx.Context) *Docs {
	return &Docs{Context: context}
}

// CreateDocsPage creates a documentation page.
func (d *Docs) CreateDocsPage() Node {
	return Div(Text(d.Context.Localizer.Localize(consts.BundleDocsIntro)))
}
