package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./telegram-bot.db"

type Storage struct {
	database *sql.DB
}

func New() Storage {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return Storage{
		database: db,
	}
}
