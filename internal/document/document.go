package document

import (
	"base/internal/document/topbar"
	"base/pkg/i18n"
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Document struct {
	Localizer *i18n.Localizer
}

func NewDocument(localizer *i18n.Localizer) *Document {
	return &Document{Localizer: localizer}
}

// CreateDocument creates an HTML document with the given title and content.
// The document includes the necessary scripts and stylesheets for the page to function properly.
// The content area is a slot, which can be updated dynamically using HTMX.
func (p *Document) CreateDocument(info *Info, content Node) Node {
	return HTML5(HTML5Props{
		Title:    info.Title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/public/styles/generated/output.css"), Type("text/css"), Defer()),
			Script(Src("/public/scripts/document.js"), Defer()),
			Script(Src("/public/scripts/lib/htmx.min.js"), Defer()),
			Script(Src("/public/scripts/lib/alpinejs.min.js"), Defer()),
		},
		Body: []Node{
			Body(
				Div(Class("h-[80vh]"), topbar.NewNavBar(p.Localizer, info.Path).CreateNavbar(),
					Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
						Div(Class("h-full rounded-lg bg-base-200"),
							Main(Class("h-full p-5"), ID("content"),
								p.CreatePartial(content)))))),
		},
	})
}

// CreatePartial creates a container for dynamic content.
// It is used when creating the document and for a snippet when swapping content with HTMX.
// The included script offers the minimally required functionality for one snippet.
func (p *Document) CreatePartial(content Node) Node {
	return Div(content,
		Script(Src("/public/scripts/partial.js"), Defer()))
}
