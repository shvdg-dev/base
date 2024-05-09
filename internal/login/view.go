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
		Div(Class("pt-4 flex flex-col space-y-3"),
			loginForm(),
			registerLink(),
			resetLink()))
}

func loginForm() Node {
	return FormEl(Attr("hx-post", "/login"),
		Div(Class("flex flex-col space-y-2"),
			mailField(),
			passwordField(),
			loginButton()))
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
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"), Text("Login"))
}

func registerLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Not yet an account?"),
			Text(" "),
			Label(Text("Register"), Class("link link-info cursor-pointer")),
			Text(".")))
}

func resetLink() Node {
	return A(hx.PushURL("true"), hx.Target("#content"),
		Div(Class("italic"),
			Text("Forgot your password?"),
			Text(" "),
			Label(Text("Reset your password"), Class("link link-info cursor-pointer")),
			Text(".")))
}
