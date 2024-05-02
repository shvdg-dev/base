package topbar

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

type pageLink struct {
	Path string
	Name string
}

func NavBar() Node {
	return Div(Class("h-15 navbar bg-base-200"),
		navBarStart(),
		navBarCenter(),
		navBarEnd(),
	)
}

func navBarStart() Node {
	return Div(Class("navbar-start pl-12"),
		A(hx.Get("/home"), hx.PushURL("true"), hx.Target("#content"),
			Class("btn btn-ghost text-xl normal-case"),
			Label(Class("text-base-content"), Text("BACK")),
			Label(Class("text-primary"), Text("2")),
			Label(Class("text-base-content"), Text("BASICS")),
			Label(Class("text-base-content"), icons.Globe())),
	)
}

func navBarCenter() Node {
	return Div(Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map([]pageLink{
				{Path: "/home", Name: "Home"},
				{Path: "/docs", Name: "Docs"},
			}, func(pl pageLink) Node { return navLink(pl.Path, pl.Name) }))))
}

func navBarEnd() Node {
	return Div(Class("navbar-end pr-12"), Options())
}

func navLink(href, name string) Node {
	return Li(A(hx.Get(href), hx.PushURL("true"), hx.Target("#content"), Text(name)))
}