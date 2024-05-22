package components

import (
	ctx "base/internal/context"
	"base/internal/document/info"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Navbar struct {
	Context  *ctx.Context
	NavItems []NavItem
}

type NavItem struct {
	Path     string
	Name     string
	IsActive bool
}

func NewNavBar(context *ctx.Context) *Navbar {
	navbar := &Navbar{Context: context}
	return navbar
}

func (n *Navbar) NewNavItem(path, name string) NavItem {
	return NavItem{
		Path:     path,
		Name:     name,
		IsActive: false,
	}
}

func (n *Navbar) CreateNavbar(info *info.Info) Node {
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.CreateEmblem(),
		n.CreateNavItems(info),
		n.CreateOptions(info),
	)
}

func (n *Navbar) CreateEmblem() Node {
	return Div(Class("navbar-start pl-12 cursor-pointer"),
		A(hx.Get("/home"), Class("text-xl flex space-x-2"),
			Div(Class("text-base-content"), Text("BACK")),
			Div(Class("text-primary"), Text("2")),
			Div(Class("text-base-content"), Text("BASICS"))),
	)
}

func (n *Navbar) CreateNavItems(info *info.Info) Node {
	n.NavItems = []NavItem{
		n.NewNavItem("/home", n.Context.Localizer.Localize("Home")),
		n.NewNavItem("/docs", n.Context.Localizer.Localize("Docs"))}
	n.setActiveLink(info.Path)
	return Div(ID("menu-items"), hx.SwapOOB("true"), Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map(n.NavItems, func(navItem NavItem) Node {
				return navItem.CreateNavItem()
			}))))
}

func (n *Navbar) setActiveLink(currentPath string) {
	for index, link := range n.NavItems {
		if link.Path == currentPath {
			n.NavItems[index].IsActive = true
		}
	}
}

func (n *NavItem) CreateNavItem() Node {
	return Li(A(hx.Get(n.Path), Text(n.Name),
		Classes{"border-b-4": true, "border-transparent": !n.IsActive, "border-primary": n.IsActive}),
	)
}

func (n *Navbar) CreateOptions(info *info.Info) Node {
	return Div(Class("navbar-end pr-12"),
		Div(Class("dropdown dropdown-end"), hx.PushURL("true"), hx.Target("#content"),
			Button(Class("btn btn-circle border-primary"), icons.User()),
			Div(Class("dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52 space-y-1"),
				TabIndex("0"),
				Button(Class("btn btn-sm"), icons.User(), Text("Profile"), hx.Get("/profile")),
				Button(Class("btn btn-sm"), icons.Cog(), Text("Settings"), hx.Get("/settings")),
				n.CreateInOutButton(info))))
}

func (n *Navbar) CreateInOutButton(info *info.Info) Node {
	base := Group([]Node{ID("login-logout"), hx.SwapOOB("true"), Class("btn btn-sm")})
	if info.IsAuthenticated {
		return Button(base, icons.LogOut(), Text("Logout"), hx.Get("/logout"))
	} else {
		return Button(base, icons.LogIn(), Text("Login"), hx.Get("/login"))
	}
}
