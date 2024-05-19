package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	"base/internal/files"
	"base/internal/home"
	"base/internal/login"
	"base/pkg/database"
	"base/pkg/environment"
	"base/pkg/i18n"
	"base/pkg/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const (
	port               = "127.0.0.1:3000"
	databaseUrl        = "DATABASE_URL"
	englishTranslation = "resources/translations/en.toml"
)

func main() {
	context := ctx.NewContext(
		initializeDatabase(),
		initializeLogger(),
		initializeTranslator())
	router := initializeRouter(context)
	startServer(router)
}

func initializeDatabase() *database.Connection {
	URL := environment.GetValueAsString(databaseUrl)
	return database.NewConnection(URL)
}

func initializeLogger() *logger.Logger {
	return logger.NewLogger()
}

func initializeTranslator() *i18n.Translator {
	trans := i18n.NewTranslator()
	trans.AddTranslations(englishTranslation)
	return trans
}

func initializeRouter(context *ctx.Context) chi.Router {
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
