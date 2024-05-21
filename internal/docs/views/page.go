package views

import (
	ctx "base/internal/context"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type Docs struct {
	Context *ctx.Context
}

func NewDocs(context *ctx.Context) *Docs {
	return &Docs{Context: context}
}

func (d *Docs) CreateDocsPage() Node {
	return Div(Text(d.Context.Localizer.Localize("DocsIntroduction")))
}
