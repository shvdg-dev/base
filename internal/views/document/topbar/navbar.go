package topbar

import (
	"base/internal/models"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type pageLink struct {
	Path string
	Name string
}

type navbar struct {
	currentPage models.Page
}

func NavBar(currentPage models.Page) Node {
	navbar := navbar{currentPage: currentPage}
	return navbar.navBar()
}

func (n *navbar) navBar() Node {
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.navBarStart(),
		n.navBarCenter(),
		n.navBarEnd(),
	)
}

func (n *navbar) navBarStart() Node {
	return Div(Class("navbar-start pl-12"),
		A(hx.Get("/home"),
			Class("btn btn-ghost text-xl normal-case"),
			Label(Class("text-base-content"), Text("BACK")),
			Label(Class("text-primary"), Text("2")),
			Label(Class("text-base-content"), Text("BASICS")),
			Label(Class("text-base-content"), icons.Globe())),
	)
}

func (n *navbar) navBarCenter() Node {
	return Div(Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map([]pageLink{
				{Path: "/home", Name: "Home"},
				{Path: "/docs", Name: "Docs"},
			}, func(pl pageLink) Node {
				return navLink(pl.Path, pl.Name, n.currentPage.Path == pl.Path)
			}))))
}

func (n *navbar) navBarEnd() Node {
	return Div(Class("navbar-end pr-12"), Options())
}

func navLink(href, name string, isActive bool) Node {
	return Li(A(hx.Get(href), Text(name)), Classes{"underline": isActive})
}
