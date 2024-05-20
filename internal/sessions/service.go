package sessions

import (
	"base/pkg/database"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
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

func (s *Service) Store(key string, value interface{}, request *http.Request) {
	s.Manager.Put(request.Context(), key, value)
}

func (s *Service) Get(key string, request *http.Request) interface{} {
	return s.Manager.Get(request.Context(), key)
}

func (s *Service) IsAuthenticated(request *http.Request) bool {
	isAuthenticated, ok := s.Get("isAuthenticated", request).(bool)
	if !ok {
		return false
	}
	return isAuthenticated
}
