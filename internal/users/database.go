package users

import (
	"base/pkg/environment"
	_ "github.com/lib/pq"

	"base/internal/app"
	"log"
)

func init() {
	createUsersTable()
	insertAdmin()
}

// createUsersTable creates the User table.
func createUsersTable() {
	_, err := app.Connections.Database.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// insertAdmin inserts the default admin.
func insertAdmin() {
	email := environment.GetValueAsString("ADMIN_INITIAL_EMAIL")
	password := environment.GetValueAsString("ADMIN_INITIAL_PASSWORD")
	InsertUser(email, password)
}

// InsertUser inserts a user with the provided email and password.
func InsertUser(email, password string) {
	_, err := app.Connections.Database.Exec(insertUserQuery, email, password, "dummy")
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect verifies whether the password is set for the user with the email.
func IsPasswordCorrect(email, password string) bool {
	var foundPassword string
	err := app.Connections.Database.QueryRow(selectUserPasswordQuery, email).Scan(&foundPassword)
	if err != nil {
		log.Fatal(err)
	}
	return foundPassword == password
}
