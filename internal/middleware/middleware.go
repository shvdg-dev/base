package middleware

import (
	ctx "base/internal/context"
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	vi "base/internal/views"
	"net/http"
	"strings"
)

type Middleware struct {
	Context  *ctx.Context
	Views    *vi.Views
	Renderer *rend.Renderer
}

func NewMiddleware(context *ctx.Context, views *vi.Views, renderer *rend.Renderer) *Middleware {
	return &Middleware{Context: context, Views: views, Renderer: renderer}
}

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		isAuthenticated := m.Context.Sessions.IsAuthenticated(request)
		if !IsResourceAccessible(request.URL.Path) && (!isAuthenticated) {
			writer.WriteHeader(http.StatusUnauthorized)
			info := m.Context.Informer.NewInfo(request, doc.WithTitle("401 - Unauthorized, whoops!"))
			m.Renderer.Render(writer, request, info, m.Views.Error.CreateAuthenticationRequiredPage())
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func IsResourceAccessible(path string) bool {
	return strings.HasPrefix(path, "/public") ||
		strings.HasPrefix(path, "/login")
}
