package layout

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"tab-collector/models"
)

func Page(p models.Page) Node {
	return HTML5(HTML5Props{
		Title:    p.Title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/static/styles/output.css"), Type("text/css")),
			Script(Src("/static/scripts/htmx.min.js")),
		},
		Body: []Node{
			Body(Class("h-[80vh]"),
				Container(Main(ID("content"), Partial(p.Content))),
			),
		},
	})
}

func Container(children ...Node) Node {
	return Div(Class("h-full pt-5 pb-5 pl-20 pr-20"), Div(Class("h-full rounded-lg bg-base-200"), Div(Class("p-5"), Group(children))))
}

func Partial(children ...Node) Node {
	return Div(Class("h-full"), Group(children))
}
