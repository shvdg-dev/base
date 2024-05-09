package document

import (
	"base/internal/document/topbar"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Document creates an HTML document with the given title and content.
// The document includes the necessary scripts and stylesheets for the page to function properly.
// The content area is passed as a slot, which can be updated dynamically using HTMX.
//
// Parameters:
//   - title: the title of the document
//   - content: the main content of the document
//
// Returns:
//   - The generated HTML document as a Node.
func Document(page Page, content Node) Node {
	return HTML5(HTML5Props{
		Title:    page.Title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/public/styles/generated/output.css"), Type("text/css"), Defer()),
			Script(Src("/public/scripts/document.js"), Defer()),
			Script(Src("/public/scripts/websockets.js"), Defer()),
			Script(Src("/public/scripts/lib/htmx.min.js"), Defer()),
			Script(Src("/public/scripts/lib/alpinejs.min.js"), Defer()),
		},
		Body: []Node{
			Body(
				Div(Class("h-[80vh]"), topbar.NavBar(page.Path),
					Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
						Div(Class("h-full rounded-lg bg-base-200"),
							Main(Class("h-full p-5"), ID("content"),
								Slot(content)))))),
		},
	})
}

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