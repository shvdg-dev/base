package views

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	info "base/internal/document/info"
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Login is used for views regarding logging in.
type Login struct {
	Context *ctx.Context
	Info    *info.Info
}

// NewLogin creates a new Login.
func NewLogin(context *ctx.Context) *Login {
	return &Login{Context: context}
}

// CreateLoginPage creates a login page.
func (l *Login) CreateLoginPage(info *info.Info) Node {
	l.Info = info
	return Div(
		Header(Text(l.Context.Localizer.Localize(consts.BundleWelcome))),
		Div(Class("pt-4 flex flex-col space-y-3"),
			l.CreateLoginForm(),
			l.CreateRequestErrors(),
			l.CreateAccountRegisterLink(),
			l.CreatePasswordResetLink()))
}

// CreateLoginForm Creates the login form.
func (l *Login) CreateLoginForm() Node {
	return FormEl(hx.Post(consts.PathLogin), hx.Target("#content"),
		Div(Class("flex flex-col space-y-2"),
			l.CreateMailField(),
			l.CreatePasswordField(),
			l.CreateLoginButton()))
}

// CreateMailField creates the email field.
func (l *Login) CreateMailField() Node {
	return Label(components.Classes{"border-red-500": l.Info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Mail(components.Classes{"text-red-500": l.Info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder(l.Context.Localizer.Localize(consts.BundleEmail)),
		),
	)
}

// CreatePasswordField creates the password field.
func (l *Login) CreatePasswordField() Node {
	return Label(components.Classes{"border-red-500": l.Info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Key(components.Classes{"text-red-500": l.Info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder(l.Context.Localizer.Localize(consts.BundlePassword)),
		),
	)
}

// CreateRequestErrors creates an overview of all errors encountered during a request.
func (l *Login) CreateRequestErrors() Node {
	var errors []Node
	for _, errMessage := range l.Info.Errors {
		errors = append(errors, Div(Role("alert"), Class("alert alert-error w-60"), Span(Text(errMessage))))
	}
	return Group(errors)
}

// CreateLoginButton creates the login button.
func (l *Login) CreateLoginButton() Node {
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"), Text(l.Context.Localizer.Localize(consts.BundleLogin)))
}

// CreateAccountRegisterLink creates the register link.
func (l *Login) CreateAccountRegisterLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text(l.Context.Localizer.Localize(consts.BundleRegisterQuestion)),
			Text(" "),
			Label(Text(l.Context.Localizer.Localize(consts.BundleRegister)), Class("link link-info cursor-pointer")),
			Text(".")))
}

// CreatePasswordResetLink creates the reset link.
func (l *Login) CreatePasswordResetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text(l.Context.Localizer.Localize(consts.BundleForgotPasswordQuestion)),
			Text(" "),
			Label(Text(l.Context.Localizer.Localize(consts.BundleResetPassword)), Class("link link-info cursor-pointer")),
			Text(".")))
}
