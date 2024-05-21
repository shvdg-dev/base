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
	Context *ctx.Context
	Views   *vi.Views
}

func NewMiddleware(context *ctx.Context, views *vi.Views) *Middleware {
	return &Middleware{Context: context, Views: views}
}

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		isAuthenticated := m.Context.Sessions.IsAuthenticated(request)
		if !IsResourceAccessible(request.URL.Path) && (!isAuthenticated) {
			writer.WriteHeader(http.StatusUnauthorized)
			info := doc.NewInfo(request, doc.WithTitle("401 - Unauthorized, whoops!"))
			rend.GetRenderer().Render(writer, request, info, m.Views.Error.CreateAuthenticationRequiredPage())
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func IsResourceAccessible(path string) bool {
	return strings.HasPrefix(path, "/public") ||
		strings.HasPrefix(path, "/login")
}
