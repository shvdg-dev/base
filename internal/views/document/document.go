package document

import (
	"base/internal/views/components"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Document(title string, content Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/public/styles/output.css"), Type("text/css")),
			Script(Src("/public/scripts/htmx.min.js")),
		},
		Body: []Node{
			Body(Class("h-[80vh]"),
				components.NavBar(),
				Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
					Div(Class("h-full rounded-lg bg-base-200"),
						Main(ID("content"), Class("h-full p-5"), content)))),
		},
	})
}
