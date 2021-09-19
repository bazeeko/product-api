package database

import (
	"database/sql"
	"simple-api/config"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(config config.Config) (db *sql.DB, err error) {
	db, err = sql.Open(config.Database.Driver, config.Database.FilePath)
	if err != nil {
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		name TEXT UNIQUE NOT NULL,
		description TEXT,
		price REAL NOT NULL
	);`)
	return
}
