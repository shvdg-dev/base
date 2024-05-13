package docs

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Page() Node {
	return Div(Text("Hello Docs!"))
}
