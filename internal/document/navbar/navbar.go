package navbar

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	"base/internal/info"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Navbar represents a navigation bar.
type Navbar struct {
	Context  *ctx.Context
	NavItems []NavItem
}

// NavItem represents an item in a navigation bar.
type NavItem struct {
	Path     string
	Name     string
	IsActive bool
}

// NewNavBar creates a new instance of Navbar.
func NewNavBar(context *ctx.Context) *Navbar {
	navbar := &Navbar{Context: context}
	return navbar
}

// NewNavItem creates a new navigation item.
func NewNavItem(path, name string) NavItem {
	return NavItem{
		Path:     path,
		Name:     name,
		IsActive: false,
	}
}

// CreateNavbar creates a navigation bar.
func (n *Navbar) CreateNavbar(info *info.Info) Node {
	return Div(Class("h-15 navbar bg-base-200"), hx.PushURL("true"), hx.Target("#content"),
		n.CreateEmblem(),
		n.CreateNavItems(info),
		n.CreateOptions(info),
	)
}

// CreateEmblem creates an emblem for the navigation bar.
func (n *Navbar) CreateEmblem() Node {
	return Div(Class("navbar-start pl-12 cursor-pointer"),
		A(hx.Get(consts.PathHome), Class("text-xl flex space-x-2"),
			Div(Class("text-base-content"), Text("BACK")),
			Div(Class("text-primary"), Text("2")),
			Div(Class("text-base-content"), Text("BASICS"))),
	)
}

// CreateNavItems creates the navigation items for the navigation bar.
func (n *Navbar) CreateNavItems(info *info.Info) Node {
	n.NavItems = []NavItem{
		NewNavItem(consts.PathHome, n.Context.Localizer.Localize(consts.BundleHome)),
		NewNavItem(consts.PathDocs, n.Context.Localizer.Localize(consts.BundleDocs))}
	n.setActiveNavItem(info.Path)
	return Div(ID("menu-items"), hx.SwapOOB("true"), Class("navbar-center"),
		Ul(Class("menu menu-horizontal px-1"),
			Group(Map(n.NavItems, func(navItem NavItem) Node {
				return navItem.CreateNavItem()
			}))))
}

// setActiveNavItem makes the navigation item active if the current path is equal to the path of the navigation item.
func (n *Navbar) setActiveNavItem(currentPath string) {
	for index, link := range n.NavItems {
		if link.Path == currentPath {
			n.NavItems[index].IsActive = true
		}
	}
}

// CreateNavItem creates a navigation item node for a NavItem struct.
func (n *NavItem) CreateNavItem() Node {
	return Li(A(hx.Get(n.Path), Text(n.Name),
		Classes{"border-b-4": true, "border-transparent": !n.IsActive, "border-primary": n.IsActive}),
	)
}

// CreateOptions creates the options section of the navigation bar.
func (n *Navbar) CreateOptions(info *info.Info) Node {
	return Div(Class("navbar-end pr-12"),
		Div(Class("dropdown dropdown-end"), hx.PushURL("true"), hx.Target("#content"),
			Button(Class("btn btn-circle border-primary"), icons.User()),
			Div(Class("dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52 space-y-1"),
				TabIndex("0"),
				Button(Class("btn btn-sm"), icons.User(), Text(n.Context.Localizer.Localize(consts.BundleProfile)), hx.Get(consts.BundleHome)),
				Button(Class("btn btn-sm"), icons.Cog(), Text(n.Context.Localizer.Localize(consts.BundleSettings)), hx.Get(consts.BundleHome)),
				n.CreateInOutButton(info))))
}

// CreateInOutButton creates the login/logout button as part of the options in the navigation bar.
func (n *Navbar) CreateInOutButton(info *info.Info) Node {
	base := Group([]Node{ID("login-logout"), hx.SwapOOB("true"), Class("btn btn-sm")})
	if info.IsAuthenticated {
		return Button(base, icons.LogOut(), Text(n.Context.Localizer.Localize(consts.BundleLogout)), hx.Get(consts.PathLogout))
	} else {
		return Button(base, icons.LogIn(), Text(n.Context.Localizer.Localize(consts.BundleLogin)), hx.Get(consts.PathLogin))
	}
}
