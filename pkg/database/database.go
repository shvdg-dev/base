package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Connection struct {
	DB *sql.DB
}

func NewConnection(URL string) *Connection {
	database, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return &Connection{DB: database}
}
