package storage

import (
	"distributed_calculator/internal/config"
	model2 "distributed_calculator/internal/model"
	"distributed_calculator/internal/model/expression"
)

type Repository interface {
	CreateExpression(*expression.Expression) error
	UpdateExpression(*expression.Expression) error
	ReadExpressions() ([]*expression.Expression, error)
	ReadAllExpressionsWithStatus(expression.Status) ([]*expression.Expression, error)
	ReadExpression(int) (*expression.Expression, error)

	CreateOperation(*model2.OperationWithDuration) error
	ReadOperations() ([]*model2.OperationWithDuration, error)
	ReadOperation(operationType model2.OperationType) (*model2.OperationWithDuration, error)
	UpdateOperation(*model2.OperationWithDuration) error
	SeedOperation(config *config.Config) error

	Init(config *config.Config) error
}

type RepositoryQuickAccess interface {
	Init(config *config.Config) error
	StoreIdempotencyToken(string, string, *model2.ResponseData) error
	RetrieveIdempotencyToken(string, string) (*model2.ResponseData, error)
}
