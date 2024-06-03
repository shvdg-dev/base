package views

import (
	consts "base/internal/constants"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// CreateHomePage creates a home page.
func (h *Home) CreateHomePage() Node {
	return Div(Text(h.Context.Localizer.Localize(consts.BundleHomeIntro)))
}
