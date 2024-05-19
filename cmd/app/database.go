package main

import (
	ctx "base/internal/context"
	"base/pkg/database"
	"base/pkg/environment"
)

func initializeDatabase() *database.Connection {
	URL := environment.GetValueAsString(databaseUrl)
	return database.NewConnection(URL)
}

func populateDatabase(context *ctx.Context) {
	context.UserService.CreateUsersTable()
	insertAdmin(context)
}

func insertAdmin(context *ctx.Context) {
	email := environment.GetValueAsString(adminInitialEmailKey)
	password := environment.GetValueAsString(adminInitialPasswordKey)
	context.UserService.InsertUser(email, password)
}
