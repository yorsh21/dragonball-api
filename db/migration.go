package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Migration struct {
	DB *sql.DB
}

func NewMigration(db *sql.DB) Migration {
	return Migration{
		DB: db,
	}
}

func (m Migration) Character() {
	sql := `
    CREATE TABLE IF NOT EXISTS characters (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        ki TEXT
    );
    `
	_, err := m.DB.Exec(sql)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sql)
	}

	fmt.Println("Characters table created or already exist.")
}
