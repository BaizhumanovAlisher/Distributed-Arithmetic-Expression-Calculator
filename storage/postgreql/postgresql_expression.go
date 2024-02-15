package postgreql

import (
	"database/sql"
	"distributed_calculator/model/expression"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresqlDB struct {
	db *sql.DB
}

func (s *PostgresqlDB) ReadAllExpressionsWithStatus(status expression.Status) ([]*expression.Expression, error) {
	rows, err := s.db.Query(`SELECT id, expression, answer, status, created_at, completed_at FROM expressions WHERE status = $1`, status)

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

func (s *PostgresqlDB) CreateExpression(expression *expression.Expression) error {
	stmt, err := s.db.Prepare(`INSERT INTO expressions (expression, answer, status, created_at, completed_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(expression.Expression, expression.Answer, expression.Status, expression.CreatedAt, expression.CompletedAt).Scan(&expression.Id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (s *PostgresqlDB) ReadAllExpressions() ([]*expression.Expression, error) {
	rows, err := s.db.Query(`SELECT id, expression, answer, status, created_at, completed_at FROM expressions`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all expressions: %w", err)
	}
	defer rows.Close()

	var expressions []*expression.Expression
	for rows.Next() {
		expr := new(expression.Expression)
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
	row := s.db.QueryRow(`SELECT id, expression, answer, status, created_at, completed_at FROM expressions WHERE id = $1`, id)

	expr := new(expression.Expression)
	err := row.Scan(&expr.Id, &expr.Expression, &expr.Answer, &expr.Status, &expr.CreatedAt, &expr.CompletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("failed to scan row into expression: %w", err)
	}

	return expr, nil
}
