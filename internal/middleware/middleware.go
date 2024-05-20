package middleware

import (
	ctx "base/internal/context"
	doc "base/internal/document/info"
	err "base/internal/error"
	"net/http"
	"strings"
)

type Middleware struct {
	Context *ctx.Context
}

func NewMiddleware(context *ctx.Context) *Middleware {
	return &Middleware{Context: context}
}

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		isAuthenticated := m.Context.Sessions.IsAuthenticated(request)
		if !IsResourceAccessible(request.URL.Path) && (!isAuthenticated) {
			writer.WriteHeader(http.StatusUnauthorized)
			m.Context.Renderer.Render(
				doc.NewInfo(request, doc.WithTitle("401 - Unauthorized, whoops!")),
				err.NewPage(m.Context).CreateAuthenticationRequiredPage(),
				writer, request)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func IsResourceAccessible(path string) bool {
	return strings.HasPrefix(path, "/public") ||
		strings.HasPrefix(path, "/login")
}
