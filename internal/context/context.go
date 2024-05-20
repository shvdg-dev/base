package context

import (
	rend "base/internal/renderer"
	sess "base/internal/sessions"
	"base/internal/users"
	"base/pkg/database"
	"base/pkg/i18n"
	"base/pkg/logger"
)

type Context struct {
	Database  *database.Connection
	Logger    *logger.Logger
	Localizer *i18n.Localizer
	Sessions  *sess.Service
	Renderer  *rend.Renderer
	Users     *users.Service
}

func NewContext(database *database.Connection, logger *logger.Logger, localizer *i18n.Localizer) *Context {
	sessions := sess.NewService(database)
	renderer := rend.NewRenderer(localizer, sessions)
	return &Context{
		Database:  database,
		Logger:    logger,
		Localizer: localizer,
		Sessions:  sessions,
		Renderer:  renderer,
		Users:     users.NewService(database)}
}
