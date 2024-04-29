package views

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Home() Node {
	return Div(Text("Hello World!"))
}
