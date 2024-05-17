package home

import (
	"base/internal/app"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Page() Node {
	return Div(Text(app.I18n.GetText("Home")))
}
