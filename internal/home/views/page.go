package views

import (
	ctx "base/internal/context"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

type Home struct {
	Context *ctx.Context
}

func NewHome(context *ctx.Context) *Home {
	return &Home{Context: context}
}

func (h *Home) CreateHomePage() Node {
	return Div(Text(h.Context.Localizer.Localize("HomeIntroduction")))
}
