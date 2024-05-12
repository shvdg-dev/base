package docs

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func page() Node {
	return Div(Text("Hello Docs!"))
}
