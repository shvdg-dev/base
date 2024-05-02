package document

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Wrapper(content Node) Node {
	return Div(content,
		Link(Rel("stylesheet"), Href("/public/styles/output.css"), Type("text/css"), Defer()),
		Script(Src("/public/scripts/htmx.min.js"), Defer()),
		Script(Src("/public/scripts/document.js"), Defer()))
}
