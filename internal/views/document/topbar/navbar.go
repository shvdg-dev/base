package topbar

import (
	"fmt"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type navbar struct {
	currentPath string
	pageLinks   []pageLink
}

type pageLink struct {
	Path     string
	Name     string
	IsActive bool
}

func NavBar(currentPath string) Node {
	navbar := navbar{
		currentPath: currentPath,
		pageLinks: []pageLink{
			{Path: "/home", Name: "Home"},
			{Path: "/docs", Name: "Docs"},
		}}
	return navbar.navBar()
}

func (n *navbar) navBar() Node {
	n.setActiveLink()
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.navBarStart(),
		n.navBarCenter(),
		n.navBarEnd(),
	)
}

func (n *navbar) setActiveLink() {
	for index, link := range n.pageLinks {
		if link.Path == n.currentPath {
			n.pageLinks[index].IsActive = true
		}
	}
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
			Group(Map(n.pageLinks, func(pageLink pageLink) Node {
				return navLink(pageLink)
			}))))
}

func (n *navbar) navBarEnd() Node {
	return Div(Class("navbar-end pr-12"), Options())
}

func navLink(pageLink pageLink) Node {
	return Li(A(
		ID(fmt.Sprintf("menu-item-%s", pageLink.Path)),
		hx.Get(pageLink.Path),
		Text(pageLink.Name),
		Classes{"underline": pageLink.IsActive}),
	)
}
