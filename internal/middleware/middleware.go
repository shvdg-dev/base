package middleware

import (
	ctx "base/internal/context"
	"net/http"
)

type Middleware struct {
	Context *ctx.Context
}

func NewMiddleware(context *ctx.Context) *Middleware {
	return &Middleware{Context: context}
}

func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		isAuthenticated, ok := (m.Context.Sessions.Get("isAuthenticated", request)).(bool)
		if request.URL.Path != "/login" && (!isAuthenticated || !ok) {
			http.Redirect(writer, request, "/login", 302)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func (m *Middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
	})
}
