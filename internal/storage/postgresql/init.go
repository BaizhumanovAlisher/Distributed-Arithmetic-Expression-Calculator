package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"internal/helpers"
)

type PostgresqlDB struct {
	db *sql.DB
}

func NewPostgresql(cfg *helpers.Config) (*PostgresqlDB, error) {
	db, err := sql.Open("postgres", cfg.Storage.StoragePath)

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
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS expressions (
    id SERIAL PRIMARY KEY,
    expression TEXT NOT NULL,
    answer VARCHAR NULL,
    status VARCHAR,
    created_at timestamp NOT NULL,
    completed_at timestamp NULL,
    user_id BIGINT REFERENCES users NOT NULL 
);

CREATE TABLE IF NOT EXISTS operations (
    id SERIAL PRIMARY KEY,
    operation_kind VARCHAR UNIQUE NOT NULL ,
    duration_in_sec INT NOT NULL 
);
`

	if _, err := s.db.Exec(q); err != nil {
		return err
	}

	if cfg.Operation.PermissionToSeed {
		err := s.SeedOperation(cfg)
		if err != nil {
			return err
		}
	}

	return nil
}
