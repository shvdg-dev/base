package home

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func View() Node {
	return Div(Text("Hello Home!"))
}
