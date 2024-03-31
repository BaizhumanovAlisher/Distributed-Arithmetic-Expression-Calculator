package postgresql

import (
	"database/sql"
	"fmt"
	"internal/helpers"
)

func Postgresql(cfg *helpers.Config) (*PostgresqlDB, error) {
	storage := cfg.Storage

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		storage.User,
		storage.Password,
		storage.Host,
		storage.Port,
		storage.DBName,
		storage.SSLMode,
	)

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	postgresql := &PostgresqlDB{db: db}
	err = postgresql.Init(cfg)

	if err != nil {
		return nil, err
	}

	return postgresql, nil
}

func (s *PostgresqlDB) Init(cfg *helpers.Config) error {
	q := `
CREATE TABLE IF NOT EXISTS expressions (
    id SERIAL PRIMARY KEY,
    expression TEXT,
    answer VARCHAR,
    status VARCHAR,
    created_at timestamp,
    completed_at timestamp NULL 
);

CREATE TABLE IF NOT EXISTS operations (
    id SERIAL PRIMARY KEY,
    operation_kind VARCHAR UNIQUE,
    duration_in_sec INT
);
`

	if _, err := s.db.Exec(q); err != nil {
		return err
	}

	err := s.SeedOperation(cfg)
	if err != nil {
		return err
	}
	return err
}
