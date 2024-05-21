package views

import (
	ctx "base/internal/context"
	. "base/internal/docs/views"
	. "base/internal/document/components"
	. "base/internal/error"
	. "base/internal/home/views"
	. "base/internal/login/views"
)

type Views struct {
	Navbar *Navbar
	Home   *Home
	Docs   *Docs
	Login  *Login
	Error  *Error
}

func NewViews(context *ctx.Context) *Views {
	return &Views{
		Navbar: NewNavBar(context),
		Home:   NewHome(context),
		Docs:   NewDocs(context),
		Login:  NewLogin(context),
		Error:  NewError(context),
	}
}
