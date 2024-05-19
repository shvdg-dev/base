package home

import (
	ctx "base/internal/context"
	"github.com/go-chi/chi/v5"
)

type Home struct {
	context *ctx.Context
}

func NewHome(context *ctx.Context) *Home {
	return &Home{context: context}
}

func (h *Home) SetupRouter(router chi.Router) {
	router.Get("/", h.HandleHomePage)
	router.Get("/home", h.HandleHomePage)
}
