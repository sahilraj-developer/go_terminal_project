package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "tasks.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		completed BOOLEAN
	);`

	_, err = DB.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}
}