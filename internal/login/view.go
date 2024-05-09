package login

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func View() Node {
	return Div(
		Header(Text("Welcome")),
		Div(Class("pt-4"), mailField()),
		Div(Class("pt-2"), passwordField()),
		Div(Class("pt-3"), loginButton()),
		Div(Class("pt-4"), registerLink()),
		Div(Class("pt-2"), resetLink()))
}

func mailField() Node {
	return Label(Class("input input-bordered flex items-center gap-2 max-w-md"),
		icons.Mail(),
		Input(Class("grow"),
			Attr("type", "email"),
			Attr("name", "email"),
			Attr("autocomplete", "email"),
			Placeholder("E-mail"),
		),
	)
}

func passwordField() Node {
	return Label(Class("input input-bordered flex items-center gap-2 max-w-md"),
		icons.Key(),
		Input(Class("grow"),
			Attr("type", "password"),
			Attr("name", "password"),
			Attr("autocomplete", "current-password"),
			Placeholder("Password"),
		),
	)
}

func loginButton() Node {
	return Button(Class("btn btn-primary"), hx.PushURL("false"), hx.Target("#content"),
		Text("Login"))
}

func registerLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Not yet an account?"),
			Text(" "),
			Label(Text("Register"), Class("underline text-info")),
			Text(".")))
}

func resetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Forgot your password?"),
			Text(" "),
			Label(Text("Reset your password"), Class("underline text-info")),
			Text(".")))
}
