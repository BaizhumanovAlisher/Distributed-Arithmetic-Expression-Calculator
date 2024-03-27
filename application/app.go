package application

import (
	"distributed_calculator/storage"
	"distributed_calculator/storage/postgreql"
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
	Repo   *postgreql.PostgresqlDB
	Redis  *storage.RedisDB
}
