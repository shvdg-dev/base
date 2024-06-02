package users

import (
	"base/pkg/database"
	"base/pkg/utils"
	_ "github.com/lib/pq"

	"log"
)

// Service is for managing users.
type Service struct {
	Database *database.Connection
}

// NewService creates a new instance of the Service struct.
func NewService(database *database.Connection) *Service {
	return &Service{Database: database}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (u *Service) CreateUsersTable() {
	_, err := u.Database.DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertUser inserts a new user into the users table.
func (u *Service) InsertUser(email, plainPassword string) {
	hashedPassword, _ := utils.HashPassword(plainPassword)
	_, err := u.Database.DB.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect checks if the given password is correct for the user with the given email.
func (u *Service) IsPasswordCorrect(email, plainPassword string) bool {
	if email == "" || plainPassword == "" {
		return false
	}
	var foundHashedPassword string
	err := u.Database.DB.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return utils.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
