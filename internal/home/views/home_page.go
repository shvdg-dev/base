package views

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Home is used for views regarding the home page.
type Home struct {
	Context *ctx.Context
}

// NewHome initializes a new instance of Home.
func NewHome(context *ctx.Context) *Home {
	return &Home{Context: context}
}

// CreateHomePage creates a home page.
func (h *Home) CreateHomePage() Node {
	return Div(Text(h.Context.Localizer.Localize(consts.BundleHomeIntro)))
}
