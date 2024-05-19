package users

import (
	"base/pkg/database"
	"base/pkg/utils"
	_ "github.com/lib/pq"

	"log"
)

type Service struct {
	Database *database.Connection
}

func NewService(database *database.Connection) *Service {
	return &Service{Database: database}
}

// CreateUsersTable creates the User table.
func (u *Service) CreateUsersTable() {
	_, err := u.Database.DB.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertUser inserts a user with the provided email and password.
func (u *Service) InsertUser(email, plainPassword string) {
	hashedPassword, _ := utils.HashPassword(plainPassword)
	_, err := u.Database.DB.Exec(insertUserQuery, email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// IsPasswordCorrect verifies whether the password is set for the user with the email.
func (u *Service) IsPasswordCorrect(email, plainPassword string) bool {
	var foundHashedPassword string
	err := u.Database.DB.QueryRow(selectUserPasswordQuery, email).Scan(&foundHashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return utils.ArePasswordsEqual(plainPassword, foundHashedPassword)
}
