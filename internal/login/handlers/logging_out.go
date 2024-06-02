package handlers

import (
	"base/internal/constants"
	info2 "base/internal/info"
	"net/http"
)

// LoggingOut handles the process of logging out a user.
func (l *Login) LoggingOut(writer http.ResponseWriter, request *http.Request) {
	l.Context.Sessions.Store(constants.ValueIsAuthenticated, false, request)
	info := l.Context.Informer.NewInfo(request, info2.WithTitle(l.Context.Localizer.Localize(constants.BundleNotAuthenticatedTitle)), info2.WithPath(constants.PathHome))
	l.Renderer.Render(writer, request, info,
		l.Views.Error.CreateNotAuthenticatedPage(),
		l.Views.Navbar.CreateInOutButton(info))
}
