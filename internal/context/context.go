package context

import (
	"base/internal/users"
	"base/pkg/database"
	"base/pkg/i18n"
	"base/pkg/logger"
)

type Context struct {
	Database   *database.Connection
	Logger     *logger.Logger
	Translator *i18n.Translator

	//Services
	UserService *users.Service
}

func NewContext(database *database.Connection, logger *logger.Logger, translator *i18n.Translator) *Context {
	return &Context{
		Database:    database,
		Logger:      logger,
		Translator:  translator,
		UserService: users.NewService(database)}
}
