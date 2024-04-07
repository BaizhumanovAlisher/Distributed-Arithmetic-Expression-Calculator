package storage

import (
	"context"
	"internal/helpers"
	model2 "internal/model"
	"internal/model/expression"
)

// Repository todo: divide interfaces
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

	CreateUser(ctx context.Context, user *model2.User) (int64, error)
	ReadUserByName(ctx context.Context, name string) (user *model2.User, err error)

	Init(config *helpers.Config) error
}
