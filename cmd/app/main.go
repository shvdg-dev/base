package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	midware "base/internal/middleware"
	"base/pkg/i18n"
	logr "base/pkg/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	context := ctx.NewContext(initDatabase(), initLogger(), initLocalizer())
	router := initRouter(context)
	prepareDatabase(context)
	startServer(router)
}

func initLogger() *logr.Logger {
	return logr.NewLogger()
}

func initLocalizer() *i18n.Localizer {
	trans := i18n.NewLocalizer()
	trans.AddLocalization(englishTranslation)
	return trans
}

func initRouter(context *ctx.Context) chi.Router {
	router := chi.NewRouter()
	initMiddleware(router, context)
	files.SetupRouter(router)
	home.NewHome(context).SetupRouter(router)
	docs.NewDocs(context).SetupRouter(router)
	login.NewLogin(context).SetupRouter(router)
	return router
}

func initMiddleware(router chi.Router, context *ctx.Context) *midware.Middleware {
	middleware := midware.NewMiddleware(context)
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
