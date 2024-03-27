package postgreql

import (
	"database/sql"
	"distributed_calculator/config"
	"distributed_calculator/model"
	"errors"
	"fmt"
)

func (s *PostgresqlDB) CreateOperation(operation *model.OperationWithDuration) error {
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

func (s *PostgresqlDB) ReadAllOperations() ([]*model.OperationWithDuration, error) {
	rows, err := s.db.Query(`SELECT operation_kind, duration_in_sec FROM operations`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all operations: %w", err)
	}
	defer rows.Close()

	var operations []*model.OperationWithDuration
	for rows.Next() {
		op := new(model.OperationWithDuration)
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

func (s *PostgresqlDB) UpdateOperation(operation *model.OperationWithDuration) error {
	stmt, err := s.db.Prepare(`UPDATE operations SET duration_in_sec = $1 WHERE operation_kind = $2`)
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

func (s *PostgresqlDB) SeedOperation(cfg *config.Config) error {
	operationsInDatabase, err := s.ReadAllOperations()

	if err != nil {
		return err
	}

	if len(operationsInDatabase) == cfg.Operation.CountOperation {
		return nil
	}

	operations := []*model.OperationWithDuration{
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

func (s *PostgresqlDB) ReadOperation(operationType model.OperationType) (*model.OperationWithDuration, error) {
	row := s.db.QueryRow(`SELECT duration_in_sec FROM operations WHERE operation_kind = $1`, operationType)

	operationWithDuration := new(model.OperationWithDuration)
	err := row.Scan(&operationWithDuration.DurationInSecond)
	operationWithDuration.OperationKind = operationType

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, fmt.Errorf("failed to scan row into expression: %w", err)
	}

	return operationWithDuration, nil
}
