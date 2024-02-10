package storage

import (
	"distributed_calculator/config"
	"distributed_calculator/model"
)

type Repository interface {
	CreateExpression(*model.Expression) error
	ReadAllExpressions() ([]*model.Expression, error)
	ReadExpression(int) (*model.Expression, error)

	CreateOperation(*model.Operation) error
	ReadAllOperations() ([]*model.Operation, error)
	UpdateOperation(*model.Operation) error
	SeedOperation(config *config.Config) error

	Init(config *config.Config) error
}

type RepositoryQuickAccess interface {
	Init(config *config.Config) error
	StoreIdempotencyToken(string, string, *model.ResponseData) error
	RetrieveIdempotencyToken(string, string) (*model.ResponseData, error)
}
