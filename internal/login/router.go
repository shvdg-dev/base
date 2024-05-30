package login

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	rend "base/internal/renderer"
	"base/internal/views"
	"github.com/go-chi/chi/v5"
)

// Login is used for routing and handlers regarding logging in.
type Login struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewLogin creates a new Login instance.
func NewLogin(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Login {
	return &Login{Context: context, Views: views, Renderer: renderer}
}

// SetupRouter sets up the router for the Login struct.
func (l *Login) SetupRouter(router chi.Router) {
	router.Get(consts.PathLogin, l.HandleLoginPage)
	router.Post(consts.PathLogin, l.HandleLoggingIn)
	router.Get(consts.PathLogout, l.HandleLoggingOut)
}
