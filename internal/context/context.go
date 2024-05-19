package context

import (
	"base/internal/renderer"
	"base/internal/sessions"
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
	Sessions  *sessions.Service
	Users     *users.Service
}

func NewContext(database *database.Connection, logger *logger.Logger, localizer *i18n.Localizer) *Context {
	return &Context{
		Database:  database,
		Logger:    logger,
		Localizer: localizer,
		Renderer:  renderer.NewRenderer(localizer),
		Sessions:  sessions.NewService(database),
		Users:     users.NewService(database)}
}
