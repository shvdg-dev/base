package views

import (
	ctx "base/internal/context"
	. "base/internal/docs/views"
	. "base/internal/document/navbar"
	. "base/internal/error/views"
	. "base/internal/home/views"
	. "base/internal/login/views"
)

// Views represents a collection of views, made accessible throughout the app.
type Views struct {
	Navbar *Navbar
	Home   *Home
	Docs   *Docs
	Login  *Login
	Error  *Error
}

// NewViews creates a new instance of the Views structure.
func NewViews(context *ctx.Context) *Views {
	return &Views{
		Navbar: NewNavBar(context),
		Home:   NewHome(context),
		Docs:   NewDocs(context),
		Login:  NewLogin(context),
		Error:  NewError(context),
	}
}
