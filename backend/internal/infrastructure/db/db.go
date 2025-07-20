package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dbURL string) *sql.DB {
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *sql.DB) error {
	content, err := os.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(content)); err != nil {
		return err
	}
	return nil
}
