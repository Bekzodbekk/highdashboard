package postgres

import (
	"database/sql"
	load "debt-service/internal/pkg/config"
	"debt-service/storage"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(cfg load.Config) (*sql.DB, error) {
	target := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, cfg.Postgres.User, cfg.Postgres.Password)
	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Queries(db *sql.DB) storage.Queries {
	return *storage.New(db)
}
