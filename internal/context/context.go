package context

import (
	"base/internal/renderer"
	"base/internal/users"
	"base/pkg/database"
	"base/pkg/i18n"
	"base/pkg/logger"
)

type Context struct {
	Database  *database.Connection
	Logger    *logger.Logger
	Localizer *i18n.Localizer
	Renderer  *renderer.Renderer

	//Services
	UserService *users.Service
}

func NewContext(database *database.Connection, logger *logger.Logger, localizer *i18n.Localizer, renderer *renderer.Renderer) *Context {
	return &Context{
		Database:    database,
		Logger:      logger,
		Localizer:   localizer,
		Renderer:    renderer,
		UserService: users.NewService(database)}
}
