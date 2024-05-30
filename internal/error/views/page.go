package views

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

// Error is used for views regarding errors.
type Error struct {
	Context *ctx.Context
}

// NewError creates a new instance of Error.
func NewError(context *ctx.Context) *Error {
	return &Error{Context: context}
}

// CreateNotAuthenticatedPage creates a page signifying that authentication is required.
func (e *Error) CreateNotAuthenticatedPage() Node {
	return Div(
		Div(Class("flex"), icons.ShieldX(Class("mr-2 text-yellow-500")), Text(e.Context.Localizer.Localize(consts.BundleAuthenticationRequired))),
		Div(Class("pt-5"),
			Button(Class("btn btn-primary btn-md"), icons.LogIn(), Text(e.Context.Localizer.Localize(consts.BundleNavigateToLogin)),
				hx.PushURL("true"), hx.Target("#content"), hx.Get(consts.PathLogin))))
}

// CreatePageNotFound creates a page signifying that a particular page could not be found.
func (e *Error) CreatePageNotFound() Node {
	return Div(
		Div(Class("flex"), icons.ShieldQuestion(Class("mr-2 text-yellow-500")), Text(e.Context.Localizer.Localize(consts.BundlePageNotFound))),
		Div(Class("pt-5"),
			Button(Class("btn btn-primary btn-md"), icons.Home(), Text(e.Context.Localizer.Localize(consts.BundleNavigateToHome)),
				hx.PushURL("true"), hx.Target("#content"), hx.Get(consts.PathHome))))
}
