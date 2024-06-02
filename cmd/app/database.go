package main

import (
	ctx "base/internal/context"
	"base/pkg/base/database"
	"base/pkg/base/environment"
)

// createDatabase initializes the database connection by retrieving the database URL from the environment.
func createDatabase() *database.Connection {
	URL := environment.GetValueAsString(databaseUrlKey)
	return database.NewConnection(URL)
}

// prepareDatabase prepares the database by creating tables and inserting data.
func prepareDatabase(context *ctx.Context) {
	context.Users.CreateUsersTable()
	context.Sessions.CreateSessionsTable()
	insertAdmin(context)
}

// insertAdmin inserts an administrator user into the database.
func insertAdmin(context *ctx.Context) {
	email := environment.GetValueAsString(adminInitialEmailKey)
	password := environment.GetValueAsString(adminInitialPasswordKey)
	context.Users.InsertUser(email, password)
}
