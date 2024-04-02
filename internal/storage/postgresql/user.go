package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"internal/helpers"
	"internal/model"
)

func (s *PostgresqlDB) CreateUser(user *model.User) (int64, error) {
	stmt, err := s.db.Prepare(`
INSERT INTO users (name, hashed_password, created_at) 
VALUES ($1, $2, now() AT TIME ZONE 'UTC') RETURNING id
`)

	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Name, user.HashedPassword).Scan(&user.Id)
	if err != nil {
		var repoErr *pq.Error
		ok := errors.As(err, &repoErr)

		//todo: test
		if ok && repoErr.Code == "23505" {
			return 0, helpers.UsernameExistErr
		}

		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	return user.Id, nil
}

func (s *PostgresqlDB) ReadUserByName(name string) (*model.User, error) {
	row := s.db.QueryRow(`
SELECT id, hashed_password, created_at
FROM users WHERE name = $1
`, name)

	user := new(model.User)

	err := row.Scan(&user.Id, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, helpers.NoRowsErr
		}
		return nil, fmt.Errorf("failed to scan row into expression: %w", err)
	}
	user.Name = name

	return user, nil
}
