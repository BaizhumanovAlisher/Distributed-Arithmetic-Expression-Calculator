package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"internal/helpers"
	"internal/model/expression"
)

// ReadAllExpressionsWithStatus Only for expression_manager
func (s *PostgresqlDB) ReadAllExpressionsWithStatus(status expression.Status) ([]*expression.Expression, error) {
	rows, err := s.db.Query(`SELECT id, expression, answer, status, created_at, completed_at, user_id FROM expressions WHERE status = $1`, status)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query all expressions: %w", err)
	}
	defer rows.Close()

	var expressions []*expression.Expression
	for rows.Next() {
		expr := new(expression.Expression)
		err := rows.Scan(&expr.Id, &expr.Expression, &expr.Answer, &expr.Status, &expr.CreatedAt, &expr.CompletedAt, &expr.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into expression: %w", err)
		}
		expressions = append(expressions, expr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration over rows: %w", err)
	}

	return expressions, nil
}

// UpdateExpression Only for expression_manager
func (s *PostgresqlDB) UpdateExpression(e *expression.Expression) error {
	stmt, err := s.db.Prepare(`
UPDATE expressions SET 
answer = $1,
status = $2,
completed_at = $3
WHERE id = $4`)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		e.Answer,
		e.Status,
		e.CompletedAt,
		e.Id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (s *PostgresqlDB) CreateExpression(expr *expression.Expression) (int, error) {
	stmt, err := s.db.Prepare(`
INSERT INTO expressions (expression, status, created_at, user_id) 
VALUES ($1, $2, now() AT TIME ZONE 'UTC', $3) RETURNING id
`)

	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(expr.Expression, expr.Status, expr.UserId).Scan(&expr.Id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	return expr.Id, nil
}

func (s *PostgresqlDB) ReadExpressions(userId int64) ([]*expression.Expression, error) {
	rows, err := s.db.Query(`
SELECT id, expression, answer, status, created_at, completed_at 
FROM expressions WHERE user_id = $1
`, userId)

	if err != nil {
		return nil, fmt.Errorf("failed to query all expressions: %w", err)
	}
	defer rows.Close()

	var expressions []*expression.Expression
	for rows.Next() {
		expr := new(expression.Expression)
		expr.UserId = userId

		err := rows.Scan(&expr.Id, &expr.Expression, &expr.Answer, &expr.Status, &expr.CreatedAt, &expr.CompletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into expression: %w", err)
		}
		expressions = append(expressions, expr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration over rows: %w", err)
	}

	return expressions, nil
}

func (s *PostgresqlDB) ReadExpression(id int) (*expression.Expression, error) {
	row := s.db.QueryRow(`
SELECT expression, answer, status, created_at, completed_at, user_id 
FROM expressions WHERE id = $1
`, id)

	expr := new(expression.Expression)

	expr.Id = id

	err := row.Scan(&expr.Expression, &expr.Answer, &expr.Status, &expr.CreatedAt, &expr.CompletedAt, &expr.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, helpers.NoRowsErr
		}
		return nil, fmt.Errorf("failed to scan row into expression: %w", err)
	}

	return expr, nil
}
