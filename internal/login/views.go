package login

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type ViewOption func(*View)

type View struct {
	Errors []error
}

func WithErrors(e []error) ViewOption {
	return func(p *View) {
		p.Errors = e
	}
}

func NewView(opts ...ViewOption) *View {
	page := &View{}
	for _, opt := range opts {
		opt(page)
	}
	return page
}

func (v *View) page() Node {
	return Div(
		Header(Text("Welcome")),
		Div(Class("pt-4 flex flex-col space-y-3"),
			v.loginForm(),
			v.registerLink(),
			v.resetLink()))
}

func (v *View) loginForm() Node {
	return FormEl(hx.Post("/login"), hx.Target("#content"),
		Div(Class("flex flex-col space-y-2"),
			v.mailField(),
			v.passwordField(),
			v.loginButton()))
}

func (v *View) mailField() Node {
	return Label(components.Classes{"border-red-500": len(v.Errors) > 0, "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Mail(),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder("E-mail"),
		),
	)
}

func (v *View) passwordField() Node {
	return Label(components.Classes{"border-red-500": len(v.Errors) > 0, "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Key(),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder("Password"),
		),
	)
}

func (v *View) loginButton() Node {
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"), Text("Login"))
}

func (v *View) registerLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Not yet an account?"),
			Text(" "),
			Label(Text("Register"), Class("link link-info cursor-pointer")),
			Text(".")))
}

func (v *View) resetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Forgot your password?"),
			Text(" "),
			Label(Text("Reset your password"), Class("link link-info cursor-pointer")),
			Text(".")))
}
