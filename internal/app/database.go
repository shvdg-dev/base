package app

import (
	"base/pkg/environment"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Database *sql.DB

func init() {
	Database = connectDatabase()
}

func connectDatabase() *sql.DB {
	URL := environment.GetValueAsString("DATABASE_URL")

	database, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return database
}
