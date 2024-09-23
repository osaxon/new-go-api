package database

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
)

//go:embed schema.sql
var ddl string

func NewDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:my-db.db")
	if err != nil {
		return nil, err
	}
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
