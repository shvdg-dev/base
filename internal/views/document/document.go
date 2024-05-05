package document

import (
	"base/internal/views/document/topbar"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Document(title string, content Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/public/styles/generated/output.css"), Type("text/css"), Defer()),
			Script(Src("/public/scripts/document.js"), Defer()),
			Script(Src("/public/scripts/lib/htmx.min.js"), Defer()),
			Script(Src("/public/scripts/lib/alpinejs.min.js"), Defer()),
		},
		Body: []Node{
			Body(
				Div(Class("h-[80vh]"), topbar.NavBar(),
					Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
						Div(Class("h-full rounded-lg bg-base-200"),
							Main(Class("h-full p-5"), ID("content"), Slot(content)))))),
		},
	})
}
