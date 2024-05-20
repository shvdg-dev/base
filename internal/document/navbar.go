package document

import (
	ctx "base/internal/context"
	"base/internal/document/info"
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

func (d *Document) NewNavBar(info *info.Info, request *http.Request) *Navbar {
	navbar := &Navbar{Context: d.Context, Info: info, Request: request}
	navbar.NavItems = []NavItem{
		navbar.NewNavItem("/home", navbar.Context.Localizer.Localize("Home")),
		navbar.NewNavItem("/docs", navbar.Context.Localizer.Localize("Docs"))}
	return navbar
}

type NavItem struct {
	Path     string
	Name     string
	IsActive bool
}

func (n *Navbar) NewNavItem(path, name string) NavItem {
	return NavItem{
		Path:     path,
		Name:     name,
		IsActive: false,
	}
}

func (n *Navbar) CreateNavbar() Node {
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.CreateNavBarStart(),
		n.CreateNavBarCenter(),
		n.CreateNavBarEnd(),
	)
}

func (n *Navbar) CreateNavBarStart() Node {
	return Div(Class("navbar-start pl-12 cursor-pointer"),
		A(hx.Get("/home"), Class("text-xl flex space-x-2"),
			Div(Class("text-base-content"), Text("BACK")),
			Div(Class("text-primary"), Text("2")),
			Div(Class("text-base-content"), Text("BASICS"))),
	)
}

func (n *Navbar) CreateNavBarCenter() Node {
	n.setActiveLink()
	return Div(ID("navbar-center"), hx.SwapOOB("true"), Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map(n.NavItems, func(navItem NavItem) Node {
				return navItem.CreateNavItem()
			}))))
}

func (n *Navbar) setActiveLink() {
	for index, link := range n.NavItems {
		if link.Path == n.Info.Path {
			n.NavItems[index].IsActive = true
		}
	}
}

func (n *Navbar) CreateNavBarEnd() Node {
	return Div(ID("navbar-end"), hx.SwapOOB("true"), Class("navbar-end pr-12"), n.CreateOptions())
}

func (n *NavItem) CreateNavItem() Node {
	return Li(A(hx.Get(n.Path), Text(n.Name),
		Classes{"border-b-4": true, "border-transparent": !n.IsActive, "border-primary": n.IsActive}),
	)
}
