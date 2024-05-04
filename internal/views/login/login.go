package login

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Page() Node {
	return Div(Header(Text("Welcome")),
		Input(Placeholder("E-mail")),
		Input(Placeholder("Password")),
		Button(Text("Login")))
}
