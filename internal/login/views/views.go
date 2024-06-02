package views

import ctx "base/internal/context"

// Login is used for views regarding logging in.
type Login struct {
	Context *ctx.Context
}

// NewLogin creates a new Login.
func NewLogin(context *ctx.Context) *Login {
	return &Login{Context: context}
}
