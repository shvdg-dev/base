package document

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Document(title string, content Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/static/styles/output.css"), Type("text/css")),
			Script(Src("/static/scripts/htmx.min.js")),
		},
		Body: []Node{
			Body(Class("h-[80vh]"),
				Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
					Div(Class("h-full rounded-lg bg-base-200"),
						Div(Class("p-5"),
							Main(ID("component"), Div(Class("h-full"), content)))))),
		},
	})
}
