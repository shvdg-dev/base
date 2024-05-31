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
}

// NewLogin creates a new Login.
func NewLogin(context *ctx.Context) *Login {
	return &Login{Context: context}
}

// CreateLoginPage creates a login page.
func (l *Login) CreateLoginPage(info *info.Info, data *Data) Node {
	return Div(
		Header(Text(l.Context.Localizer.Localize(consts.BundleWelcome))),
		Div(Class("pt-4 flex flex-col space-y-3"),
			l.CreateLoginForm(info, data),
			l.CreateRequestErrors(info),
			l.CreateAccountRegisterLink(),
			l.CreatePasswordResetLink()))
}

// CreateLoginForm Creates the login form.
func (l *Login) CreateLoginForm(info *info.Info, data *Data) Node {
	currentEmail := ""
	currentPassword := ""
	if data != nil {
		currentEmail = data.CurrentEmail
		currentPassword = data.CurrentPassword
	}
	return FormEl(hx.Post(consts.PathLogin), hx.Target("#content"),
		Div(Class("flex flex-col space-y-2"),
			l.CreateMailField(info, currentEmail),
			l.CreatePasswordField(info, currentPassword),
			l.CreateLoginButton()))
}

// CreateMailField creates the email field.
func (l *Login) CreateMailField(info *info.Info, currentEmail string) Node {
	return Label(components.Classes{"border-red-500": info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Mail(components.Classes{"text-red-500": info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder(l.Context.Localizer.Localize(consts.BundleEmail)),
			Value(currentEmail)),
	)
}

// CreatePasswordField creates the password field.
func (l *Login) CreatePasswordField(info *info.Info, currentPassword string) Node {
	return Label(components.Classes{"border-red-500": info.HasErrors(), "input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
		icons.Key(components.Classes{"text-red-500": info.HasErrors(), "stroke-current": true}),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder(l.Context.Localizer.Localize(consts.BundlePassword)),
			Value(currentPassword)),
	)
}

// CreateRequestErrors creates an overview of all errors encountered during a request.
func (l *Login) CreateRequestErrors(info *info.Info) Node {
	var errors []Node
	for _, errMessage := range info.Errors {
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
