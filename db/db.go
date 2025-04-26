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
		panic("Could not connect to database")
	}
	
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS User (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Email TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(err)
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS Event (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Description TEXT NOT NULL,
		Location TEXT NOT NULL,
		DateTime DATETIME NOT NULL,
		UserId INTEGER NOT NULL,
		FOREIGN KEY(UserId) REFERENCES User(Id) ON DELETE CASCADE
	)
	`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic(err)
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS Registration (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		UserId INTEGER NOT NULL,
		EventId INTEGER NOT NULL,
		FOREIGN KEY(UserId) REFERENCES User(Id) ON DELETE CASCADE,
		FOREIGN KEY(EventId) REFERENCES Event(Id) ON DELETE CASCADE
	)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic(err)
	}
}