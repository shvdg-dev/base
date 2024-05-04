package document

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Wrapper(content Node) Node {
	return Div(content,
		Script(Src("/public/scripts/document.js"), Defer()))
}
