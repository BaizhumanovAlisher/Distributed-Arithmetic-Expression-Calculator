package postgreql

import (
	"database/sql"
	"distributed_calculator/config"
)

func Postgresql(cfg *config.Config) (*PostgresqlDB, error) {
	//conn := fmt.Sprintf("user=%s dbname=%s password='%s' host=%s port=%s sslmode=%s", cfg.User, cfg.DBName, cfg.Storage.Password, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := sql.Open("postgres", cfg.URL)

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

func (s *PostgresqlDB) Init(cfg *config.Config) error {
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
