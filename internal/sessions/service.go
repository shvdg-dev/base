package sessions

import (
	"base/pkg/database"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"log"
)

type Service struct {
	Database *database.Connection
	Manager  *scs.SessionManager
}

func NewService(database *database.Connection) *Service {
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(database.DB)
	return &Service{Database: database, Manager: sessionManager}
}

func (s *Service) CreateSessionsTable() {
	_, err := s.Database.DB.Exec(createSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Database.DB.Exec(createSessionExpiryIndexQuery)
	if err != nil {
		log.Fatal(err)
	}
}
