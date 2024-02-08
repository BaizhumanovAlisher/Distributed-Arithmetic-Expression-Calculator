package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"orchestrator/config"
	"orchestrator/model"

	_ "github.com/lib/pq"
)

type Sqlite struct {
	db *sql.DB
}

func Postgresql(cfg *config.Config) (*Sqlite, error) {
	conn := fmt.Sprintf("user=%s dbname=%s password='%s' host=%s port=%s sslmode=%s", cfg.User, cfg.DBName, cfg.Password, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	sqlite := &Sqlite{db: db}
	err = sqlite.Init(cfg)

	if err != nil {
		return nil, err
	}

	return sqlite, nil
}

func (s *Sqlite) CreateExpression(expression *model.Expression) error {
	stmt, err := s.db.Prepare(`INSERT INTO expressions (expression, answer, status, created_at, completed_at) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(expression.Expression, expression.Answer, expression.Status, expression.CreatedAt, expression.CompletedAt)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}
	expression.Id = int(lastID)

	return nil
}

func (s *Sqlite) ReadAllExpressions() ([]*model.Expression, error) {
	rows, err := s.db.Query(`SELECT id, expression, answer, status, created_at, completed_at FROM expressions`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all expressions: %w", err)
	}
	defer rows.Close()

	var expressions []*model.Expression
	for rows.Next() {
		expr := new(model.Expression)
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

func (s *Sqlite) ReadExpression(id int) (*model.Expression, error) {
	row := s.db.QueryRow(`SELECT id, expression, answer, status, created_at, completed_at FROM expressions WHERE id = $1`, id)

	expr := new(model.Expression)
	err := row.Scan(&expr.Id, &expr.Expression, &expr.Answer, &expr.Status, &expr.CreatedAt, &expr.CompletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no expression found with id %d: %w", id, err)
		}
		return nil, fmt.Errorf("failed to scan row into expression: %w", err)
	}

	return expr, nil
}

func (s *Sqlite) CreateOperation(operation *model.Operation) error {
	stmt, err := s.db.Prepare(`INSERT INTO operations (operation_kind, duration_in_sec) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(operation.OperationKind, operation.DurationInSecond)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (s *Sqlite) ReadAllOperations() ([]*model.Operation, error) {
	rows, err := s.db.Query(`SELECT operation_kind, duration_in_sec FROM operations`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all operations: %w", err)
	}
	defer rows.Close()

	var operations []*model.Operation
	for rows.Next() {
		op := new(model.Operation)
		err := rows.Scan(&op.OperationKind, &op.DurationInSecond)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into operation: %w", err)
		}
		operations = append(operations, op)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration over rows: %w", err)
	}

	return operations, nil
}

func (s *Sqlite) UpdateOperation(operation *model.Operation) error {
	stmt, err := s.db.Prepare(`UPDATE operations SET duration_in_sec = ? WHERE operation_kind = $1`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(operation.DurationInSecond, operation.OperationKind)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func (s *Sqlite) SeedOperation(cfg *config.Config) error {
	operationsInDatabase, err := s.ReadAllOperations()

	if err != nil {
		return err
	}

	if len(operationsInDatabase) == cfg.CountOperation {
		return nil
	}

	operations := []*model.Operation{
		{OperationKind: model.Addition, DurationInSecond: cfg.DurationInSecondAddition},
		{OperationKind: model.Subtraction, DurationInSecond: cfg.DurationInSecondSubtraction},
		{OperationKind: model.Multiplication, DurationInSecond: cfg.DurationInSecondMultiplication},
		{OperationKind: model.Division, DurationInSecond: cfg.DurationInSecondDivision},
	}

	for _, operation := range operations {
		err := s.CreateOperation(operation)
		if err != nil {
			return fmt.Errorf("failed to create operation: %w", err)
		}
	}

	return nil
}

func (s *Sqlite) Init(cfg *config.Config) error {
	q := `
CREATE TABLE IF NOT EXISTS expressions (
    id SERIAL PRIMARY KEY,
    expression TEXT,
    answer VARCHAR,
    status VARCHAR,
    created_at DATE,
    completed_at DATE
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
