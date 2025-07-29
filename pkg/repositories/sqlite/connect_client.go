package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectClient(dbURI string) *sql.DB {
	db, err := sql.Open("sqlite3", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON; PRAGMA user_version = 1;`)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
