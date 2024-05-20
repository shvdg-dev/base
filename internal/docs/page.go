package docs

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func (d *Docs) Page() Node {
	return Div(Text(d.Context.Localizer.Localize("DocsIntroduction")))
}
