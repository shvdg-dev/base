package topbar

import (
	"base/pkg/i18n"
	"fmt"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Navbar struct {
	Localizer   *i18n.Localizer
	CurrentPath string
	PageLinks   []PageLink
}

func NewNavBar(localizer *i18n.Localizer, currentPath string) *Navbar {
	navbar := &Navbar{Localizer: localizer, CurrentPath: currentPath}
	navbar.PageLinks = []PageLink{
		NewPageLink("/home", navbar.Localizer.Localize("Home")),
		NewPageLink("/docs", navbar.Localizer.Localize("Docs"))}
	return navbar
}

type PageLink struct {
	Path     string
	Name     string
	IsActive bool
}

func NewPageLink(path, name string) PageLink {
	return PageLink{
		Path:     path,
		Name:     name,
		IsActive: false,
	}
}

func (n *Navbar) CreateNavbar() Node {
	n.setActiveLink()
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.navBarStart(),
		n.navBarCenter(),
		n.navBarEnd(),
	)
}

func (n *Navbar) setActiveLink() {
	for index, link := range n.PageLinks {
		if link.Path == n.CurrentPath {
			n.PageLinks[index].IsActive = true
		}
	}
}

func (n *Navbar) navBarStart() Node {
	return Div(Class("navbar-start pl-12 cursor-pointer"),
		A(hx.Get("/home"), Class("text-xl flex space-x-2"),
			Div(Class("text-base-content"), Text("BACK")),
			Div(Class("text-primary"), Text("2")),
			Div(Class("text-base-content"), Text("BASICS"))),
	)
}

func (n *Navbar) navBarCenter() Node {
	return Div(Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map(n.PageLinks, func(pageLink PageLink) Node {
				return pageLink.CreatePageLink()
			}))))
}

func (n *Navbar) navBarEnd() Node {
	return Div(Class("navbar-end pr-12"), Options())
}

func (p *PageLink) CreatePageLink() Node {
	return Li(
		A(ID(fmt.Sprintf("menu-item-%s", p.Path)),
			hx.Get(p.Path),
			Text(p.Name),
			Classes{"border-b-4": true, "border-transparent": !p.IsActive, "border-primary": p.IsActive}),
	)
}
