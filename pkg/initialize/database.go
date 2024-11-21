package initialize

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(source string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	return db, nil
}
