package storage

import (
	"internal/helpers"
	model2 "internal/model"
	"internal/model/expression"
)

type Repository interface {
	CreateExpression(*expression.Expression) (int, error)
	UpdateExpression(*expression.Expression) error
	ReadExpressions(userId int64) ([]*expression.Expression, error)
	ReadAllExpressionsWithStatus(expression.Status) ([]*expression.Expression, error)
	ReadExpression(id int) (*expression.Expression, error)

	CreateOperation(*model2.OperationWithDuration) error
	ReadOperations() ([]*model2.OperationWithDuration, error)
	ReadOperation(operationType model2.OperationType) (*model2.OperationWithDuration, error)
	UpdateOperation(*model2.OperationWithDuration) error
	SeedOperation(config *helpers.Config) error

	CreateUser(*model2.User) (int64, error)
	ReadCredential(name string) (hashedPassword string, err error)

	Init(config *helpers.Config) error
}

type RepositoryQuickAccess interface {
	Init(config *helpers.Config) error
	StoreIdempotencyToken(string, string, *model2.ResponseData) error
	RetrieveIdempotencyToken(string, string) (*model2.ResponseData, error)
}
