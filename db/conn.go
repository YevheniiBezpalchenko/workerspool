package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Connect() (*sql.DB, error) {
	dbx, err := sql.Open(
		"mysql",
		"user:password@tcp(127.0.0.1:3306)/wpool",
	)

	if err != nil {
		return nil, err
	}

	err = dbx.Ping()
	if err != nil {
		return nil, err
	}

	dbx.SetMaxIdleConns(2)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)
	dbx.SetConnMaxLifetime(10 * time.Second)

	return dbx, nil
}
