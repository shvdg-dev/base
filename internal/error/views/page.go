package views

import (
	ctx "base/internal/context"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

type Error struct {
	Context *ctx.Context
}

func NewError(context *ctx.Context) *Error {
	return &Error{Context: context}
}

func (e *Error) CreateNotAuthenticatedPage() Node {
	return Div(
		Div(Class("flex"), icons.ShieldX(Class("mr-2 text-yellow-500")), Text("You are required to be authenticated.")),
		Div(Class("pt-5"),
			Button(Class("btn btn-primary btn-md"), icons.LogIn(), Text("Go to login"),
				hx.PushURL("true"), hx.Target("#content"), hx.Get("/login"))))
}

func (e *Error) CreatePageNotFound() Node {
	return Div(
		Div(Class("flex"), icons.ShieldQuestion(Class("mr-2 text-yellow-500")), Text("This page does not even exist!?")),
		Div(Class("pt-5"),
			Button(Class("btn btn-primary btn-md"), icons.Home(), Text("Go back home"),
				hx.PushURL("true"), hx.Target("#content"), hx.Get("/home"))))
}