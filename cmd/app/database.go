package main

import (
	ctx "base/internal/context"
	"base/pkg/database"
	"base/pkg/environment"
)

func initDatabase() *database.Connection {
	URL := environment.GetValueAsString(databaseUrl)
	return database.NewConnection(URL)
}

func prepareDatabase(context *ctx.Context) {
	context.Users.CreateUsersTable()
	context.Sessions.CreateSessionsTable()
	insertAdmin(context)
}

func insertAdmin(context *ctx.Context) {
	email := environment.GetValueAsString(adminInitialEmailKey)
	password := environment.GetValueAsString(adminInitialPasswordKey)
	context.Users.InsertUser(email, password)
}
