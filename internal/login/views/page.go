package views

import (
	ctx "base/internal/context"
	info "base/internal/document/info"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Login struct {
	Context *ctx.Context
	Info    *info.Info
}

func NewLogin(context *ctx.Context) *Login {
	return &Login{Context: context}
}

func (l *Login) CreateLoginPage(info *info.Info) Node {
	l.Info = info
	return Div(
		Header(Text("Welcome")),
		Div(Class("pt-4 flex flex-col space-y-3"),
			l.loginForm(),
			l.authenticationFail(),
			l.registerLink(),
			l.resetLink()))
}

func (l *Login) loginForm() Node {
	return FormEl(hx.Post("/login"), hx.Target("#content"),
		Div(Class("flex flex-col space-y-2"),
			l.mailField(),
			l.passwordField(),
			l.loginButton()))
}

func (l *Login) mailField() Node {
	return Label(components.Classes{"border-red-500": l.Info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Mail(components.Classes{"text-red-500": l.Info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder("E-mail"),
		),
	)
}

func (l *Login) passwordField() Node {
	return Label(components.Classes{"border-red-500": l.Info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Key(components.Classes{"text-red-500": l.Info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder("Password"),
		),
	)
}

func (l *Login) authenticationFail() Node {
	var errors []Node
	for _, errMessage := range l.Info.Errors {
		errors = append(errors, Div(Role("alert"), Class("alert alert-error w-60"), Span(Text(errMessage))))
	}
	return Group(errors)
}

func (l *Login) loginButton() Node {
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"), Text("Login"))
}

func (l *Login) registerLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Not yet an account?"),
			Text(" "),
			Label(Text("Register"), Class("link link-info cursor-pointer")),
			Text(".")))
}

func (l *Login) resetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Forgot your password?"),
			Text(" "),
			Label(Text("Reset your password"), Class("link link-info cursor-pointer")),
			Text(".")))
}
