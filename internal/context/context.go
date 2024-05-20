package context

import (
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
	Sessions  *sess.Service
	Users     *users.Service
}

func NewContext(database *database.Connection, localizer *i18n.Localizer) *Context {
	return &Context{
		Database:  database,
		Localizer: localizer,
		Logger:    loggr.NewLogger(),
		Sessions:  sess.NewService(database),
		Users:     users.NewService(database)}
}
