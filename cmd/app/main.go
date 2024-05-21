package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	midware "base/internal/middleware"
	"base/internal/renderer"
	vi "base/internal/views"
	"base/pkg/i18n"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	context := ctx.NewContext(initDatabase(), initLocalizer())
	views := vi.NewViews(context)
	renderer.NewRenderer(context, views)
	router := initRouter(context, views)
	prepareDatabase(context)
	startServer(router)
}

func initLocalizer() *i18n.Localizer {
	trans := i18n.NewLocalizer()
	trans.AddLocalization(englishTranslation)
	return trans
}

func initRouter(context *ctx.Context, views *vi.Views) chi.Router {
	router := chi.NewRouter()
	initMiddleware(router, context, views)
	files.SetupRouter(router)
	home.NewHome(context, views).SetupRouter(router)
	docs.NewDocs(context, views).SetupRouter(router)
	login.NewLogin(context, views).SetupRouter(router)
	return router
}

func initMiddleware(router chi.Router, context *ctx.Context, views *vi.Views) *midware.Middleware {
	middleware := midware.NewMiddleware(context, views)
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
