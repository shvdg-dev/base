package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"tab-collector/models"
)

func Page(p models.Page) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    p.Title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/static/styles/output.css"), Type("text/css")),
			Script(Src("/static/scripts/htmx.min.js")),
		},
		Body: []g.Node{
			Body(Class("h-[80vh]"),
				Container(Main(ID("content"), Partial(p.Content))),
			),
		},
	})
}

func Container(children ...g.Node) g.Node {
	return Div(Class("h-full pt-5 pb-5 pl-20 pr-20"), Div(Class("h-full rounded-lg bg-base-200"), Div(Class("p-5"), g.Group(children))))
}

func Partial(children ...g.Node) g.Node {
	return Div(Class("h-full"), g.Group(children))
}
