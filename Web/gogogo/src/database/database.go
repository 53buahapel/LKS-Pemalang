package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := healthCheck(DB); err != nil {
		log.Fatal(err)
	}

	initData(DB)

	log.Println("Database connected")
}

func initData(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL,
		author_id INTEGER NOT NULL,
		FOREIGN KEY (author_id) REFERENCES users(id)
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func healthCheck(db *sql.DB) error {
	return db.Ping()
}
