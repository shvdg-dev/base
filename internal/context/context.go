package context

import (
	"base/internal/document/info"
	sess "base/internal/sessions"
	"base/internal/users"
	"base/pkg/database"
	"base/pkg/i18n"
	loggr "base/pkg/logger"
)

type Context struct {
	Database  *database.Connection
	Localizer *i18n.Localizer
	Logger    *loggr.Logger
	Informer  *info.Informer
	Sessions  *sess.Service
	Users     *users.Service
}

func NewContext(database *database.Connection, localizer *i18n.Localizer) *Context {
	sessions := sess.NewService(database)
	return &Context{
		Database:  database,
		Localizer: localizer,
		Logger:    loggr.NewLogger(),
		Informer:  info.NewInformant(sessions),
		Sessions:  sessions,
		Users:     users.NewService(database)}
}
