package storage

import "orchestrator/model"

type Repository interface {
	CreateExpression(*model.Expression) error
	ReadAllExpressions() ([]*model.Expression, error)
	ReadExpression(int) (*model.Expression, error)

	CreateOperation(*model.Operation) error
	ReadAllOperations() ([]*model.Operation, error)
	UpdateOperation(*model.Operation) error
	SeedOperation() error

	Init() error
}
