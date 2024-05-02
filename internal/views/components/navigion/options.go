package navigion

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func Options() Node {
	return Div(Class("dropdown dropdown-end dropdown-hover"),
		Button(Class("btn btn-circle border-primary"), icons.User()),
		Div(Class("dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52 space-y-1"),
			TabIndex("0"),
			Button(Class("btn btn-sm"), icons.User(), Text("Login")),
			Button(Class("btn btn-sm"), icons.Cog(), Text("Settings"), hx.Get(""), hx.PushURL(""), hx.Target("#content"))))
}
