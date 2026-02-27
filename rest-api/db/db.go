package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	//DB.SetConnMaxIdleTime(5)
	createTable()

}

func createTable() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events(
    id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, location TEXT NOT NULL, DateTime DATETIME NOT NULL,
    description TEXT NOT NULL, UserId TEXT NOT NULL
)`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create table")
	}

}
