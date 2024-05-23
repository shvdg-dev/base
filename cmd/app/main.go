package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	err "base/internal/error"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	midware "base/internal/middleware"
	rend "base/internal/renderer"
	vi "base/internal/views"
	"base/pkg/i18n"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	context := ctx.NewContext(initDatabase(), initLocalizer())
	views := vi.NewViews(context)
	renderer := rend.NewRenderer(context, views)
	router := initRouter(context, views, renderer)
	prepareDatabase(context)
	startServer(router)
}

func initLocalizer() *i18n.Localizer {
	trans := i18n.NewLocalizer()
	trans.AddLocalization(englishTranslation)
	return trans
}

func initRouter(context *ctx.Context, views *vi.Views, renderer *rend.Renderer) chi.Router {
	router := chi.NewRouter()
	initMiddleware(router, context, views, renderer)
	files.SetupRouter(router)
	err.NewError(context, views, renderer).SetupRouter(router)
	home.NewHome(context, views, renderer).SetupRouter(router)
	docs.NewDocs(context, views, renderer).SetupRouter(router)
	login.NewLogin(context, views, renderer).SetupRouter(router)
	return router
}

func initMiddleware(router chi.Router, context *ctx.Context, views *vi.Views, renderer *rend.Renderer) *midware.Middleware {
	middleware := midware.NewMiddleware(context, views, renderer)
	router.Use(context.Sessions.Manager.LoadAndSave)
	router.Use(middleware.Authentication)
	return middleware
}

func startServer(router chi.Router) {
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
