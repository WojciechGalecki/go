package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	maxOpenConnections = 10
	maxIdleConnections = 5
	driverName         = "sqlite3"
	dataSourceName     = "api.db"
)

var DB *sql.DB

func InitDB() {
  var err error
	DB, err = sql.Open(driverName, dataSourceName)

	if err != nil {
		panic("Couldn't connect to database")
	}

	DB.SetMaxOpenConns(maxOpenConnections)
	DB.SetMaxIdleConns(maxIdleConnections)
  createTables()
}

func createTables() {
	createEventsTableQuery := `
  CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    date_time TEXT NOT NULL,
    user_id INTEGER
  )`

  _, err := DB.Exec(createEventsTableQuery)

  if err != nil {
    panic("Couldn't create events table")
  }
}
