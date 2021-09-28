package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connection := "user=pguser port=5450 dbname=db password=pgpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
