package views

import (
	consts "base/internal/constants"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// CreateDocsPage creates a documentation page.
func (d *Docs) CreateDocsPage() Node {
	return Div(Text(d.Context.Localizer.Localize(consts.BundleDocsIntro)))
}
