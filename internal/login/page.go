package login

import (
	doc "base/internal/document"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Page struct {
	info *doc.Info
}

func (l *Login) NewPage(info *doc.Info) *Page {
	return &Page{info: info}
}

func (p *Page) CreateLoginPage() Node {
	return Div(
		Header(Text("Welcome")),
		Div(Class("pt-4 flex flex-col space-y-3"),
			p.loginForm(),
			p.authenticationFail(),
			p.registerLink(),
			p.resetLink()))
}

func (p *Page) loginForm() Node {
	return FormEl(hx.Post("/login"), hx.Target("#content"),
		Div(Class("flex flex-col space-y-2"),
			p.mailField(),
			p.passwordField(),
			p.loginButton()))
}

func (p *Page) mailField() Node {
	return Label(components.Classes{"border-red-500": p.info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Mail(),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder("E-mail"),
		),
	)
}

func (p *Page) passwordField() Node {
	return Label(components.Classes{"border-red-500": p.info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Key(),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder("Password"),
		),
	)
}

func (p *Page) authenticationFail() Node {
	var errors []Node
	for _, errMessage := range p.info.Errors {
		errors = append(errors, Div(Role("alert"), Class("alert alert-error w-60"), Span(Text(errMessage))))
	}
	return Group(errors)
}

func (p *Page) loginButton() Node {
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"), Text("Login"))
}

func (p *Page) registerLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Not yet an account?"),
			Text(" "),
			Label(Text("Register"), Class("link link-info cursor-pointer")),
			Text(".")))
}

func (p *Page) resetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Forgot your password?"),
			Text(" "),
			Label(Text("Reset your password"), Class("link link-info cursor-pointer")),
			Text(".")))
}
