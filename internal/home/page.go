package home

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func (h *Home) Page() Node {
	return Div(Text("home"))
}
