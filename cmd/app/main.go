package main

import (
	ctx "base/internal/context"
	"base/internal/docs"
	erro "base/internal/error"
	"base/internal/files"
	hand "base/internal/handlers"
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

// main is the entry point of the application.
// It initializes the context, views, renderer, router,
// prepares the database, and starts the server.
func main() {
	context := ctx.NewContext(createDatabase(), createLocalizer())
	views := vi.NewViews(context)
	renderer := rend.NewRenderer(context, views)
	handlers := hand.NewHandlers(context, views, renderer)
	router := chi.NewRouter()
	setupMiddleware(router, context, views, renderer)
	setupRouter(router, context, handlers)
	prepareDatabase(context)
	startServer(router)
}

// createLocalizer initializes a localizer instance, adding the English translation and returning the localizer.
func createLocalizer() *i18n.Localizer {
	trans := i18n.NewLocalizer()
	trans.AddLocalization(englishTranslation)
	return trans
}

// setupRouter initializes and configures the router for the application.
func setupRouter(router chi.Router, context *ctx.Context, handlers *hand.Handlers) chi.Router {
	files.SetupRouter(router)
	erro.NewError(handlers).SetupRouter(router)
	home.NewHome(handlers).SetupRouter(router)
	docs.NewDocs(handlers).SetupRouter(router)
	login.NewLogin(context, handlers).SetupRouter(router)
	return router
}

// setupMiddleware initializes the middleware using the provided context, views, and renderer.
func setupMiddleware(router chi.Router, context *ctx.Context, views *vi.Views, renderer *rend.Renderer) *midware.Middleware {
	middleware := midware.NewMiddleware(context, views, renderer)
	router.Use(context.Sessions.Manager.LoadAndSave)
	router.Use(middleware.Authentication)
	return middleware
}

// startServer starts the server using the specified router.
func startServer(router chi.Router) {
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
