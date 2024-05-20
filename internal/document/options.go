package document

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func (n *Navbar) CreateOptions() Node {
	return Div(Class("dropdown dropdown-end"), hx.PushURL("true"), hx.Target("#content"),
		Button(Class("btn btn-circle border-primary"), icons.User()),
		Div(Class("dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52 space-y-1"),
			TabIndex("0"),
			Button(Class("btn btn-sm"), icons.User(), Text("Profile"), hx.Get("/profile")),
			Button(Class("btn btn-sm"), icons.Cog(), Text("Settings"), hx.Get("/settings")),
			n.loginOrLogout()))
}

func (n *Navbar) loginOrLogout() Node {
	if n.Context.Sessions.IsAuthenticated(n.Request) {
		return Button(Class("btn btn-sm"), icons.LogOut(), Text("Logout"), hx.Get("/logout"))
	} else {
		return Button(Class("btn btn-sm"), icons.LogIn(), Text("Login"), hx.Get("/login"))
	}
}
