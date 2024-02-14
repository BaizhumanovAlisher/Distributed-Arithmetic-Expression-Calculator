package storage

import (
	"distributed_calculator/config"
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
)

type Repository interface {
	CreateExpression(*expression.Expression) error
	UpdateExpression(*expression.Expression) error
	ReadAllExpressions() ([]*expression.Expression, error)
	ReadAllExpressionsWithStatus(expression.Status) ([]*expression.Expression, error)
	ReadExpression(int) (*expression.Expression, error)

	CreateOperation(*model.OperationWithDuration) error
	ReadAllOperations() ([]*model.OperationWithDuration, error)
	ReadOperation(operationType model.OperationType) (*model.OperationWithDuration, error)
	UpdateOperation(*model.OperationWithDuration) error
	SeedOperation(config *config.Config) error

	Init(config *config.Config) error
}

type RepositoryQuickAccess interface {
	Init(config *config.Config) error
	StoreIdempotencyToken(string, string, *model.ResponseData) error
	RetrieveIdempotencyToken(string, string) (*model.ResponseData, error)
}
