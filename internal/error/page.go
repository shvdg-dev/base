package error

import (
	ctx "base/internal/context"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

type Page struct {
	Context *ctx.Context
}

func NewPage(context *ctx.Context) *Page {
	return &Page{Context: context}
}

func (e *Page) CreateAuthenticationRequiredPage() Node {
	return Div(
		Div(Class("flex"), icons.ShieldX(Class("mr-2 text-yellow-500")), Text("Authentication is required for accessing this page.")),
		Div(Class("pt-5"),
			Button(Class("btn btn-primary btn-md"), icons.LogIn(), Text("Go to login"),
				hx.PushURL("true"), hx.Target("#content"), hx.Get("/login"))))
}
