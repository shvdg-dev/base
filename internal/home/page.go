package home

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func (h *Home) CreateHomePage() Node {
	return Div(Text(h.Context.Localizer.Localize("HomeIntroduction")))
}
