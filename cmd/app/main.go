package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	rend "base/internal/renderer"
	"base/pkg/i18n"
	logr "base/pkg/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	database := initDatabase()
	logger := initLogger()
	localizer := initLocalizer()
	renderer := initRenderer(localizer)
	context := ctx.NewContext(
		database,
		logger,
		localizer,
		renderer)
	router := initRouter(context)
	populateDatabase(context)
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

func initRenderer(localizer *i18n.Localizer) *rend.Renderer {
	return rend.NewRenderer(localizer)
}

func initRouter(context *ctx.Context) chi.Router {
	router := chi.NewRouter()
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
