package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect() error {
	var err error
	var dns = "user=postgres password=CJ dbname=story host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}
	return db.Ping()
}
