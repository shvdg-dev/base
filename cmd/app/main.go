package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
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
	//TODO: Use middleware
	//middleware := midware.NewMiddleware(context)
	router := chi.NewRouter()
	//TODO: Use middleware
	//router.Use(context.Sessions.Manager.LoadAndSave, middleware.Authentication)
	files.SetupRouter(router)
	home.NewHome(context).SetupRouter(router)
	docs.NewDocs(context).SetupRouter(router)
	login.NewLogin(context).SetupRouter(router)
	return router
}

func startServer(router chi.Router) {
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
