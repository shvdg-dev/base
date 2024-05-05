package document

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

// Slot creates a container for dynamic content.
// It is used when creating the document and for a snippet when swapping content with HTMX.
// The included script offers the minimally required functionality for one snippet.
//
// Parameters:
// - content: the content to be inserted into the container.
//
// Returns:
// - The generated HTML document as a Node.
func Slot(content Node) Node {
	return Div(content,
		Script(Src("/public/scripts/slot.js"), Defer()))
}
