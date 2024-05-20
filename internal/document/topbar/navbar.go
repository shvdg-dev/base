package topbar

import (
	ctx "base/internal/context"
	"base/internal/document/info"
	"fmt"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"net/http"
)

type Navbar struct {
	Context  *ctx.Context
	Info     *info.Info
	Request  *http.Request
	NavItems []NavItem
}

func NewNavBar(context *ctx.Context, info *info.Info, request *http.Request) *Navbar {
	navbar := &Navbar{Context: context, Info: info, Request: request}
	navbar.NavItems = []NavItem{
		NewNavItem("/home", navbar.Context.Localizer.Localize("Home")),
		NewNavItem("/docs", navbar.Context.Localizer.Localize("Docs"))}
	return navbar
}

type NavItem struct {
	Path     string
	Name     string
	IsActive bool
}

func NewNavItem(path, name string) NavItem {
	return NavItem{
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
	for index, link := range n.NavItems {
		if link.Path == n.Info.Path {
			n.NavItems[index].IsActive = true
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
			Group(Map(n.NavItems, func(navItem NavItem) Node {
				return navItem.CreateNavItem()
			}))))
}

func (n *Navbar) navBarEnd() Node {
	return Div(Class("navbar-end pr-12"), n.CreateOptions())
}

func (p *NavItem) CreateNavItem() Node {
	return Li(
		A(ID(fmt.Sprintf("menu-item-%s", p.Path)),
			hx.Get(p.Path),
			Text(p.Name),
			Classes{"border-b-4": true, "border-transparent": !p.IsActive, "border-primary": p.IsActive}),
	)
}
