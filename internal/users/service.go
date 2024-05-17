package users

import (
	"base/pkg/environment"
	"base/pkg/utils"
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
	_, err := app.Database.Exec(createUsersTableQuery)
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
func InsertUser(email, plainPassword string) {
	hashedPassword, _ := utils.HashPassword(plainPassword)
	_, err := app.Database.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect verifies whether the password is set for the user with the email.
func IsPasswordCorrect(email, plainPassword string) bool {
	var foundHashedPassword string
	err := app.Database.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return utils.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
