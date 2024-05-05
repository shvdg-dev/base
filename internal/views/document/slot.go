package document

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Slot(content Node) Node {
	return Div(content,
		Script(Src("/public/scripts/slot.js"), Defer()))
}
